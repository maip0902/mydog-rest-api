package models

import (
    "github.com/globalsign/mgo/bson"
    "github.com/globalsign/mgo"

    "github.com/ant0ine/go-json-rest/rest"
    "github.com/maip0902/mydog-rest-api/mongo"
    "sync"
    "strconv"
    "fmt"
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