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
        w.WriteHeader(500)
        fmt.Println(err)
    }

    token, err := CreateToken(&user)
    if err != nil {
    // ここは考えたい
        w.WriteHeader(500)
        return
    }

    lock.RLock()
    if err := db.C("users").Insert(bson.M{"email": user.Email, "password": string(hashPass), "token": token, "verified_at": time.Now()}); err != nil {
        w.WriteHeader(500)
        return
    }
    lock.RUnlock()
}

func SignIn(w rest.ResponseWriter, r *rest.Request) {
    db = mongo.ConnectDB()
    user := models.User{}
    err := r.DecodeJsonPayload(&user)
    password := user.Password
    email := user.Email
    err = db.C("users").Find(bson.M{"email": email}).One(&user)

    if err != nil {
        fmt.Println("登録されてないユーザーです")
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(401)
        fmt.Println("パスワードが間違っています")
        return
    }
    w.WriteJson(user)
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