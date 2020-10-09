package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/globalsign/mgo"
    "net/http"
    "fmt"
    "log"
    "github.com/globalsign/mgo/bson"
    "sync"
    "strconv"
)

var db *mgo.Database

func main() {
    ConnectDB()
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    router, err := rest.MakeRouter(
        rest.Get("/code/:code", GetImageByCode),
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
   	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func ConnectDB() {
    session, _ := mgo.Dial("mongo-db:27017")
    session.SetMode(mgo.Monotonic, true)
    db = session.DB("mydog-api")
    fmt.Println("DB connect start")
}

type CodeImage struct {
    ID bson.ObjectId `bson:"_id"`
    Code  int        `bson:"code"`
    Image string     `bson:"image"`
}

var lock = sync.RWMutex{}

func GetImageByCode (w rest.ResponseWriter, r *rest.Request) {
    code, _ := strconv.Atoi(r.PathParam("code"))
    lock.RLock()
    var codeImage *CodeImage
    if err := db.C("codeImage").Find(bson.M{"code": code}).One(&codeImage); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
    w.WriteJson(codeImage)
}