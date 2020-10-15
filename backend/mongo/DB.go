package mongo

import (
    "github.com/globalsign/mgo"
    "os"
    "fmt"
)

func ConnectDB() *mgo.Database {
    session, _ := mgo.Dial(os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"))
    session.SetMode(mgo.Monotonic, true)
    db := session.DB(os.Getenv("DB_NAME"))
    fmt.Println("DB connect start")
    return db
}