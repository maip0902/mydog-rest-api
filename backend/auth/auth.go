package auth

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/globalsign/mgo"
    "github.com/maip0902/mydog-rest-api/mongo"
    "github.com/maip0902/mydog-rest-api/models"
    "github.com/globalsign/mgo/bson"
    "sync"
    "fmt"
    "log"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
)

var db *mgo.Database

var lock = sync.RWMutex{}

func SignUp(w rest.ResponseWriter, r *rest.Request) {
    db = mongo.ConnectDB()
    user := models.User{}
    err := r.DecodeJsonPayload(&user)
        err = user.CreateUserValidate()
        if err != nil {
            fmt.Printf("%v", err)
            rest.NotFound(w, r)
            return
    }

    hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password),12)
    if err != nil {
        fmt.Println(err)
    }

    token, err := CreateToken(&user)
    if err != nil {
    // ここは考えたい
        rest.NotFound(w, r)
    }

    lock.RLock()
    if err := db.C("users").Insert(bson.M{"email": user.Email, "password": string(hashPass), "token": token, "verified_at": time.Now()}); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
}

func CreateToken(user *models.User) (string, error) {
    secret := "secret"
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "email": &user.Email,
            "iss":   "__init__", // JWT の発行者が入る(文字列(__init__)は任意)
    })
    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
            log.Fatal(err)
    }

    return tokenString, nil
}