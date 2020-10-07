package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/globalsign/mgo"
    "net/http"
    "fmt"
    "log"
)

var db *mgo.Database

func main() {
    ConnectDB()
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
   		w.WriteJson(map[string]string{"Body": "Hello World!"})
   	}))
   	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func ConnectDB() {
    session, _ := mgo.Dial("mongo-db:27017")
    session.SetMode(mgo.Monotonic, true)
    db = session.DB("mydog-api")
    fmt.Println("DB connect start")
}