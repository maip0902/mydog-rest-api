package auth

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/globalsign/mgo"
    "github.com/maip0902/mydog-rest-api/mongo"
    "github.com/globalsign/mgo/bson"
    "sync"
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

var db *mgo.Database

type User struct {
    ID bson.ObjectId   `bson:"_id"`
    Name string        `bson:"name"`
    email string       `bson:"email"`
    password string    `bson:"password"`
}

var lock = sync.RWMutex{}

func SignUp(w rest.ResponseWriter, r *rest.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
//     w.Header().Set("Content-Type", "application/json")
//     w.Header().Set("Accept", "*")
//     w.Header().Set("Access-Control-Allow-Headers","*")
    db = mongo.ConnectDB()
    email := r.PathParam("email")
    password := r.PathParam("password")

    hashPass, err := bcrypt.GenerateFromPassword([]byte(password),12)
    if err != nil {
        fmt.Println(err)
    }

//     if (email == "") {
//         rest.Error(w, "emailは必須です", 500)
//     }
//
//     if (password == "") {
//         rest.Error(w, "passwordは必須です", 500)
//     }

    lock.RLock()
    if err := db.C("users").Insert(bson.M{"email": email, "password": string(hashPass)}); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()

}