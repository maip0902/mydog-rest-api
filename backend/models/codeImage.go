package models

import (
    "github.com/globalsign/mgo/bson"
    "github.com/globalsign/mgo"

    "encoding/base64"
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/maip0902/mydog-rest-api/mongo"
    "net/http"
    "sync"
    "strconv"
    "fmt"
    "os"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
)
type CodeImage struct {
    ID bson.ObjectId   `bson:"_id"`
    Code  int          `bson:"code"`
    Image string       `bson:"image"`
    Description string `bson:"description"`
}

// 読み込みと書き込みの競合解決
var lock = sync.RWMutex{}
var db *mgo.Database

func GetImageByCode (w rest.ResponseWriter, r *rest.Request) {
    code, _ := strconv.Atoi(r.PathParam("code"))

    // 読み込みlock RLock同士はブロックしない
    lock.RLock()
    db = mongo.ConnectDB()
    var codeImage *CodeImage
    if err := db.C("codeImage").Find(bson.M{"code": code}).One(&codeImage); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
    fmt.Printf("%v", codeImage)
    // HttpResponseにjson文字列を出力
    w.WriteJson(codeImage)
}

func GetImageById (w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")

    // 読み込みlock RLock同士はブロックしない
    lock.RLock()
    db = mongo.ConnectDB()
    var codeImage *CodeImage
    if err := db.C("codeImage").FindId(bson.ObjectIdHex(id)).One(&codeImage); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
    fmt.Printf("%v", codeImage)
    // HttpResponseにjson文字列を出力
    w.WriteJson(codeImage)
}

func GetAll (w rest.ResponseWriter, r *rest.Request) {

    db = mongo.ConnectDB()
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

func UpdateImage (w rest.ResponseWriter, r *rest.Request) {

    db = mongo.ConnectDB()
    var codeImage *CodeImage
    var fields = bson.M{}
    err := r.DecodeJsonPayload(&codeImage)
    data, _ := base64.StdEncoding.DecodeString(codeImage.Image)
    f, _ := os.Create("hoge.png")
    _, err = f.Write(data)
    f, err = os.Open("hoge.png")
    fmt.Println(f)
    AccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
    SecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
    MyRegion := os.Getenv("AWS_REGION")
    sess, err := session.NewSession(
     &aws.Config{
      Endpoint: aws.String("http://minio:9000"),
      Region: aws.String(MyRegion),
      S3ForcePathStyle: aws.Bool(true),
      Credentials: credentials.NewStaticCredentials(
       AccessKeyID,
       SecretAccessKey,
       "", 
      ),
    })

    uploader := s3manager.NewUploader(sess)

    up, err := uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String("code-image"),
        ACL:    aws.String("public-read"),
        Key:    aws.String("a.png"),
        Body:   f,
    })

    if err != nil {
        fmt.Println(err)
        fmt.Println("アップロードエラー")
        rest.Error(w, "予期せぬエラーが発生しました", http.StatusInternalServerError)
    }
    fmt.Println(up)
    id := codeImage.ID
    // fmt.Println(id)
    fields["description"] = codeImage.Description

    // 読み込みlock RLock同士はブロックしない
    lock.RLock()
    if err := db.C("codeImage").UpdateId(id, bson.M{"$set": fields}); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
    // fmt.Printf("%v", codeImage)
    // HttpResponseにjson文字列を出力
    w.WriteJson(codeImage)
}