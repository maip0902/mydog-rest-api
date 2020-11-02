package auth

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/globalsign/mgo"
    "github.com/maip0902/mydog-rest-api/mongo"
    "github.com/maip0902/mydog-rest-api/models"
    "github.com/globalsign/mgo/bson"
    "sync"
    "fmt"
    "golang.org/x/crypto/bcrypt"
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

    fmt.Printf("%v", email)
    fmt.Printf("%v", password)

    user := &models.User{
                Email: email,
                Password: password,
    }

    err = user.CreateUserValidate()
    fmt.Printf("%v", err)
    lock.RLock()
    if err := db.C("users").Insert(bson.M{"email": email, "password": string(hashPass)}); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
}