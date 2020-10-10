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
    "github.com/joho/godotenv"
    "os"
)

var db *mgo.Database

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
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
   	log.Fatal(http.ListenAndServe(":" + os.Getenv("DEFAULT_PORT"), api.MakeHandler()))
}

func ConnectDB() {
    session, _ := mgo.Dial(os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"))
    session.SetMode(mgo.Monotonic, true)
    db = session.DB(os.Getenv("DB_NAME"))
    fmt.Println("DB connect start")
}

type CodeImage struct {
    ID bson.ObjectId `bson:"_id"`
    Code  int        `bson:"code"`
    Image string     `bson:"image"`
}

// 読み込みと書き込みの競合解決
var lock = sync.RWMutex{}

func GetImageByCode (w rest.ResponseWriter, r *rest.Request) {
    code, _ := strconv.Atoi(r.PathParam("code"))

    // 読み込みlock RLock同士はブロックしない
    lock.RLock()
    var codeImage *CodeImage
    if err := db.C("codeImage").Find(bson.M{"code": code}).One(&codeImage); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
    // HttpResponseにjson文字列を出力
    w.WriteJson(codeImage)
}