package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/globalsign/mgo"
    "github.com/maip0902/mydog-rest-api/mongo"
    "net/http"
    "log"
    "github.com/globalsign/mgo/bson"
    "sync"
    "strconv"
    "github.com/joho/godotenv"
    "os"
    ”
)

type User struct {
    ID bson.ObjectId   `bson:"_id"`
    Name string        `bson:"name"`
    email string       `bson:"email"`
    password string    `bson:"password"`
}

func SignUp(w rest.ResponseWriter, r *rest.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")

    email, _ := strconv.Atoi(r.PathParam("email"))
    password, _ := strconv.Atoi(r.PathParam("password"))

    email == "" {
        rest.Error(w, "emailは必須です")
    }

    password == "" {
        rest.Error(w, "passwordは必須です")
    }

    lock.RLock()
    if err := db.C("users").Insert(bson.M{"email": email, "password": password}); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
}