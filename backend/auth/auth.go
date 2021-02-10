package auth

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/globalsign/mgo"
    "github.com/maip0902/mydog-rest-api/mongo"
    "github.com/maip0902/mydog-rest-api/models"
    "github.com/maip0902/mydog-rest-api/mail"
    "github.com/globalsign/mgo/bson"
    "net/http"
    "sync"
    "fmt"
    "reflect"
    "runtime"
    "os"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "time"
)

var db *mgo.Database

var lock = sync.RWMutex{}

var createTokenRetryMax = 3;

type Token struct {
    Token string
}

func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func SignUp(w rest.ResponseWriter, r *rest.Request) {
    db = mongo.ConnectDB()
    user := models.User{}
    err := r.DecodeJsonPayload(&user)

    if err != nil {
        fmt.Printf("handle: %s error: %s\n", GetFunctionName(SignUp), err.Error())
        rest.Error(w, "予期せぬエラーが発生しました", http.StatusInternalServerError)
    }

    err = user.CreateUserValidate()
    if err != nil {
        rest.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    count, err := db.C("users").Find(bson.M{"email": user.Email}).Count()
    if count > 0 {
        fmt.Println("登録済みのユーザーです")
        rest.Error(w, "登録済みのユーザーです", http.StatusBadRequest)
        return
    }

    if err != nil {
        fmt.Printf("handle: %s error: %s\n", GetFunctionName(SignUp), err.Error())
        rest.Error(w, "予期せぬエラーが発生しました", http.StatusInternalServerError)
    }

    hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password),12)
    if err != nil {
        rest.Error(w, err.Error(), 500)
        fmt.Println(err)
        return
    }

    id := bson.NewObjectId()
    user.ID = id
    emailVerifyToken, err := CreateEmailVerifyToken(&user)
    user.VerifyToken = emailVerifyToken
    // ユーザー作成
    lock.RLock()
    if err := db.C("users").Insert(bson.M{"_id": user.ID, "email": user.Email, "password": string(hashPass), "verify_token": user.VerifyToken}); err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    lock.RUnlock()
    // 仮登録メールの送信
    c, err := mail.ConnectSMTP()
    c, err = mail.SendTemporaryRegisterMail(c, user.Email, user.VerifyToken)
    err = c.Quit()
    if err != nil {
        fmt.Println(err.Error())
    }

    w.WriteJson(&user)
}

func SignIn(w rest.ResponseWriter, r *rest.Request) {
    db = mongo.ConnectDB()
    user := models.User{}
    err := r.DecodeJsonPayload(&user)
    if err != nil {
        fmt.Printf("handle: %s error: %s\n", GetFunctionName(SignIn), err.Error())
        rest.Error(w, "予期せぬエラーが発生しました", http.StatusInternalServerError)
    }

    password := user.Password
    email := user.Email
    err = db.C("users").Find(bson.M{"email": email}).One(&user)

    if err != nil {
        fmt.Println("登録されてないユーザーです")
        rest.Error(w, "登録されてないユーザーです", http.StatusNotFound)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        fmt.Println("パスワードが間違っています")
        rest.Error(w, "パスワードが間違っています", http.StatusBadRequest)
        return
    }

    // token生成を3回までretry
    var token string
    for i := 0; i < createTokenRetryMax; i++ {
        fmt.Printf("create token try:%s\n", i + 1)
        token, err = CreateToken(&user)
        if(err == nil) {
            break
        } 
        if err != nil  {
            fmt.Printf("handle: %s error: %s", GetFunctionName(CreateToken), err.Error())
            if i == createTokenRetryMax - 1 {
                rest.Error(w, err.Error(), http.StatusInternalServerError)
                fmt.Printf("handle: %s error: %s\n", GetFunctionName(SignIn), err.Error())
                return
            }
        }
    }
    fmt.Println(token)
    
    w.WriteJson(&Token{Token: token})
}

func GetAuthenticatedUser(w rest.ResponseWriter, r *rest.Request) {
    db = mongo.ConnectDB()
    token := Token{}
    err := r.DecodeJsonPayload(&token)
    if err != nil {
        fmt.Printf("handle: %s error: %s\n", GetFunctionName(GetAuthenticatedUser), err.Error())
        rest.Error(w, "予期せぬエラーが発生しました", http.StatusInternalServerError) 
    }

    getToken, err := VerifyToken(token.Token)
    // token認証はどの原因のエラーでもログイン認証やり直させる
    if getToken == nil {
        fmt.Printf("handle: verifytoken error: %s\n", err.Error())
        rest.Error(w, "セッションの有効期限が切れました", 401)
        return
    }
    claims, _ := getToken.Claims.(jwt.MapClaims)
    id := claims["jti"].(string)
    user := models.User{}
    err = db.C("users").FindId(bson.ObjectIdHex(id)).One(&user)
    if err != nil {
        fmt.Printf("handle: %s error: %s\n", GetFunctionName(GetAuthenticatedUser), err.Error())
        rest.Error(w, "予期せぬエラーが発生しました", http.StatusInternalServerError)
        return
    }
    w.WriteJson(&user)
}

func VerifyEmail(w rest.ResponseWriter, r *rest.Request) {
    // todo::既に認証メール送信していてtokenの有効期限内だったら送信できなくしたい
    t := r.URL.Query().Get("verify_token")
    fmt.Println(t)
    db = mongo.ConnectDB()
    var u *models.User
    lock.RLock()
    if err := db.C("users").Find(bson.M{"verify_token": t}).One(&u); err != nil {
        fmt.Printf("handle: %s action: mongodb %s\n", GetFunctionName(VerifyEmail), err.Error())
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
    id := u.ID
    _, err := VerifyToken(t)
    if err != nil {
        rest.Error(w, "トークンの有効期限が切れました", http.StatusUnauthorized)
    }
    lock.RLock()
    if err := db.C("users").UpdateId(id, bson.M{"verify_token": "", "verified_at": time.Now()}); err != nil {
        fmt.Printf("handle: %s action: mongodb %s\n", GetFunctionName(VerifyEmail), err.Error())
        rest.NotFound(w, r)
        return
    }
    lock.RUnlock()
    w.WriteJson(&u)
}

func CreateToken(user *models.User) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    // Create the Claims
    claims := &jwt.StandardClaims{
        ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
        Issuer:    secret,
        Id: user.ID.Hex(),
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func CreateEmailVerifyToken(user *models.User) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    // Create the Claims
    claims := &jwt.StandardClaims{
        ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
        Issuer:    secret,
        Id: user.ID.Hex(),
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error){
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
           return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(os.Getenv("JWT_SECRET")), nil
     })
     if err != nil {
        return nil, err
     }
    //  claims, ok := token.Claims.(jwt.MapClaims)
     return token, nil
}