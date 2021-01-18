package auth

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/globalsign/mgo"
    "github.com/maip0902/mydog-rest-api/mongo"
    "github.com/maip0902/mydog-rest-api/models"
    "github.com/globalsign/mgo/bson"
    "net/http"
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
        rest.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    count, err := db.C("users").Find(bson.M{"email": user.Email}).Count()
    if count > 0 {
        fmt.Println("登録済みのユーザーです")
        rest.Error(w, "登録済みのユーザーです", http.StatusBadRequest)
        return
    }

    hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password),12)
    if err != nil {
        rest.Error(w, err.Error(), 500)
        fmt.Println(err)
        return
    }

    token, err := CreateToken(&user)
    if err != nil {
        // ここは考えたい
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        fmt.Println(err)
        return
    }

    lock.RLock()
    if err := db.C("users").Insert(bson.M{"email": user.Email, "password": string(hashPass), "token": token, "verified_at": time.Now()}); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    lock.RUnlock()
    user.Token = token
    w.WriteJson(user)
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
        rest.Error(w, "登録されてないユーザーです", http.StatusNotFound)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        fmt.Println("パスワードが間違っています")
        rest.Error(w, "パスワードが間違っています", http.StatusBadRequest)
        return
    }
    w.WriteJson(user)
}

func GetAuthenticatedUser(w rest.ResponseWriter, r *rest.Request) {
    db = mongo.ConnectDB()
    user := models.User{}
    err := r.DecodeJsonPayload(&user)
    token := user.Token
    fmt.Println(token)
    err = db.C("users").Find(bson.M{"token": token}).One(&user)

    if err != nil {
        fmt.Println(err)
        rest.Error(w, "予期せぬエラーが発生しました", http.StatusInternalServerError)
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