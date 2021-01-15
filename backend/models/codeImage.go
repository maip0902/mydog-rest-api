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
    "io/ioutil"
    "fmt"
    "os"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
)
type CodeImage struct {
    ID bson.ObjectId   `bson:"_id"`
    Code  int          `bson:"code"`
    Image string       `bson:"image"`
    Description string `bson:"description"`
}

type Image struct {
    Image string   `bson:"image"`
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
    f, _ := os.Create("codeImage.png")
    _, err = f.Write(data)
    f, err = os.Open("codeImage.png")
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
    fmt.Println(codeImage.Code)
    up, err := uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String("code-image"),
        ACL:    aws.String("public-read"),
        Key:    aws.String(strconv.Itoa(codeImage.Code) + ".png"),
        Body:   f,
    })

    if err != nil {
        fmt.Println(err)
        fmt.Println("アップロードエラー")
        rest.Error(w, "予期せぬエラーが発生しました", http.StatusInternalServerError)
    }
    fmt.Println(up)
    fields["image"] = strconv.Itoa(codeImage.Code) + ".png"
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

func GetStatusImage(w rest.ResponseWriter, r *rest.Request) {
    image := r.PathParam("code")
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
    f, err := os.Create(image + ".png")
	if err != nil {
		fmt.Println(err)
	}
    downloader := s3manager.NewDownloader(sess)
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String("code-image"),
		Key:    aws.String(image + ".png"),
    })
    fmt.Println(n)
    f, err = os.Open(image + ".png")
    data, _ := ioutil.ReadAll(f)
    encodedImage := base64.StdEncoding.EncodeToString(data)
	if err != nil {
		fmt.Println(err)
    }
    w.WriteJson(&Image{encodedImage})
}