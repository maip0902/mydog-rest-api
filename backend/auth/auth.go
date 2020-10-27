package auth

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/globalsign/mgo"
    "github.com/maip0902/mydog-rest-api/mongo"
    "github.com/globalsign/mgo/bson"
    "sync"
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
    db = mongo.ConnectDB()
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("AllowedMethods", "OPTIONS")
    email := r.PathParam("email")
    password := r.PathParam("password")

//     if (email == "") {
//         rest.Error(w, "emailは必須です", 500)
//     }
//
//     if (password == "") {
//         rest.Error(w, "passwordは必須です", 500)
//     }

    lock.RLock()
    if err := db.C("users").Insert(bson.M{"email": email, "password": password}); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
}