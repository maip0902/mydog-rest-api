package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/globalsign/mgo"
    "github.com/maip0902/mydog-rest-api/mongo"
    "github.com/maip0902/mydog-rest-api/auth"
    "net/http"
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

    db = mongo.ConnectDB()
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
//     api.Use(&rest.CorsMiddleware{
//             RejectNonCorsRequests: false,
//             OriginValidator: func(origin string, request *rest.Request) bool {
//                 return origin == "http://localhost:8080"
//             },
//             AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
//             AllowedHeaders: []string{
//                 "Accept", "Content-Type", "X-Custom-Header", "Origin"},
//             AccessControlAllowCredentials: true,
//             AccessControlMaxAge:           3600,
//         })
    router, err := rest.MakeRouter(
        rest.Get("/code/:code", GetImageByCode),
        rest.Get("/code", GetAll),
        rest.Options("/signUp", auth.SignUp),
        rest.Post("/signUp", auth.SignUp),
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
   	log.Fatal(http.ListenAndServe(":" + os.Getenv("DEFAULT_PORT"), api.MakeHandler()))
}

type CodeImage struct {
    ID bson.ObjectId   `bson:"_id"`
    Code  int          `bson:"code"`
    Image string       `bson:"image"`
    Description string `bson:"description"`
}

// 読み込みと書き込みの競合解決
var lock = sync.RWMutex{}

func GetImageByCode (w rest.ResponseWriter, r *rest.Request) {
    code, _ := strconv.Atoi(r.PathParam("code"))

//     w.Header().Set("Access-Control-Allow-Origin", "*")

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

func GetAll (w rest.ResponseWriter, r *rest.Request) {
//     w.Header().Set("Access-Control-Allow-Origin", "*")

    var codeImages []*CodeImage
    // 読み込みlock RLock同士はブロックしない
    lock.RLock()
    if err := db.C("codeImage").Find(nil).All(&codeImages); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()

    // HttpResponseにjson文字列を出力
    w.WriteJson(codeImages)
}