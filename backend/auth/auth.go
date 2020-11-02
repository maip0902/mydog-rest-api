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
)

var db *mgo.Database

var lock = sync.RWMutex{}

func SignUp(w rest.ResponseWriter, r *rest.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    db = mongo.ConnectDB()

    email := r.PathParam("email")
    password := r.PathParam("password")
    hashPass, err := bcrypt.GenerateFromPassword([]byte(password),12)
    if err != nil {
        fmt.Println(err)
    }

    user := &models.User{
        Email: email,
        Password: string(hashPass),
    }
    err = user.CreateUserValidate()
//     if err != nil {
//         fmt.Printf("%v", err)
//         rest.NotFound(w, r)
//         return
//     }
    lock.RLock()
    if err := db.C("users").Insert(bson.M{"email": email, "password": string(hashPass)}); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
    token, err := CreateToken(user)
    if err != nil {
        rest.NotFound(w, r)
    }
    fmt.Println(token)
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