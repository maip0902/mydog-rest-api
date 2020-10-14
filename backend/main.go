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
//     url := CreatePreSignedUrl(&codeImage.Image)
//     &codeImage.Image = url
    // HttpResponseにjson文字列を出力
    w.WriteJson(codeImage)
}

// func CreatePreSignedUrl(ci *CodeImage) *CodeImage {
//     s, err := session.NewSession()
//
//     ak := "access_key"
//     sk := "secret_key"
//     cfg := aws.Config{
//         Credentials: credentials.NewStaticCredentials(ak, sk, ""),
//         Region: aws.String("ap-northeast-1"),
//         Endpoint: aws.String("http://127.0.0.1:9000"),
//     }
//     svc := s3.New(s, &cfg)
//     req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
//         Bucket: aws.String("code"),
//         Key:    aws.String("404.jpeg"),
//     })
//     urlStr, err := req.Presign(15 * time.Minute)
//     fmt.Printf("%v", urlStr)
//     if err != nil {
//         fmt.Println("取得失敗")
//     }
//     fmt.Println("取得成功")
//     fmt.Println("%v", urlStr)
//     ci.Image = urlStr
//     return ci
// }