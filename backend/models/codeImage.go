package models

import (
    "github.com/globalsign/mgo/bson"
    "github.com/globalsign/mgo"

    "encoding/base64"
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/maip0902/mydog-rest-api/mongo"
    "github.com/maip0902/mydog-rest-api/awssession"
    "net/http"
    "sync"
    "strconv"
    "io/ioutil"
    "fmt"
    "reflect"
    "runtime"
    "os"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
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

type UploadError struct {
    Op string
}

func (e *UploadError) Error() string {
    return fmt.Sprintf("%s", e.Op)
}

func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
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
        fmt.Printf("handle: %s action: mongodb %s\n", GetFunctionName(GetImageByCode), err.Error())
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()

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
        fmt.Printf("handle: %s action: mongodb %s\n", GetFunctionName(GetImageById), err.Error())
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
    
    // HttpResponseにjson文字列を出力
    w.WriteJson(codeImage)
}

func GetAll (w rest.ResponseWriter, r *rest.Request) {

    db = mongo.ConnectDB()
    var codeImages []*CodeImage
    // 読み込みlock RLock同士はブロックしない
    lock.RLock()
    if err := db.C("codeImage").Find(nil).All(&codeImages); err != nil {
        fmt.Printf("handle: %s action: mongodb %s\n", GetFunctionName(GetAll), err.Error())
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

    if err != nil {
        fmt.Printf("handle: %s error: %s\n", GetFunctionName(UpdateImage), err.Error())
        rest.Error(w, "予期せぬエラーが発生しました", http.StatusInternalServerError)
    }

    // 画像アップロードがあった時
    if(codeImage.Image != "") {
        err = UploadImage(codeImage)
        if err != nil {
            fmt.Printf("handle: %s action: %s\n", GetFunctionName(UploadImage), err.Error())
            rest.Error(w, "予期せぬエラーが発生しました", http.StatusInternalServerError)
        }
        fields["image"] = strconv.Itoa(codeImage.Code) + ".png"
    }
    
    id := codeImage.ID
    fields["description"] = codeImage.Description
    // 読み込みlock RLock同士はブロックしない
    lock.RLock()
    if err := db.C("codeImage").UpdateId(id, bson.M{"$set": fields}); err != nil {
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
    
    // HttpResponseにjson文字列を出力
    w.WriteJson(codeImage)
}

func UploadImage(c *CodeImage) error {
    data, _ := base64.StdEncoding.DecodeString(c.Image)
    f, _ := os.Create("codeImage.png")
    _, err := f.Write(data)
    if err != nil {
        return &UploadError{Op: "write file data"}
    }

    f, err = os.Open("codeImage.png")
    if err != nil {
        return &UploadError{Op: "open file"}
    }
    
    sess, err := awssession.StartSession()
    if err != nil {
        return &UploadError{Op: "connect s3"}
    }

    uploader := s3manager.NewUploader(sess)
    
    _, err = uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String("code-image"),
        ACL:    aws.String("public-read"),
        Key:    aws.String(strconv.Itoa(c.Code) + ".png"),
        Body:   f,
    })
    if err != nil {
        return &UploadError{Op: "upload image s3"}
    }
    
    f.Close()
    os.Remove("codeImage.png")
    return nil
}

func GetStatusImage(w rest.ResponseWriter, r *rest.Request) {
    image := r.PathParam("code")
    
    sess, err := awssession.StartSession()
    if err != nil {
        fmt.Errorf("aws error: %w", err)
        rest.Error(w, "予期せぬエラーが発生しました", http.StatusInternalServerError)
    }
    // ダウンロードする場合
    // f, err := os.Create(image + ".png")
	// if err != nil {
	// 	fmt.Println(err)
    // }
    // downloader := s3manager.NewDownloader(sess)
	// n, err := downloader.Download(f, &s3.GetObjectInput{
	// 	Bucket: aws.String("code-image"),
	// 	Key:    aws.String(image + ".png"),
    // })
    // fmt.Println(n)

    // f, err = os.Open(image + ".png")
    // data, _ := ioutil.ReadAll(f)
    // encodedImage := base64.StdEncoding.EncodeToString(data)
	// if err != nil {
	// 	fmt.Println(err)
    // }
    // w.WriteJson(&Image{encodedImage})

    // 読み込む場合
    svc := s3.New(sess)

	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("code-image"),
		Key:    aws.String(image + ".png"),
	})
	if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            case s3.ErrCodeNoSuchBucket:
                fmt.Printf("bucket %s does not exist", os.Args[0])
                rest.NotFound(w, r)
            case s3.ErrCodeNoSuchKey:
                fmt.Printf("object with key %s does not exist in bucket", os.Args[0])
                rest.NotFound(w, r)
            }
        }    
        rest.NotFound(w, r)
    }
    rc := obj.Body
    data, _ := ioutil.ReadAll(rc)
    encodedImage := base64.StdEncoding.EncodeToString(data)

    w.WriteJson(&Image{encodedImage})
}