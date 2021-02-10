package models

import (
    "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/go-ozzo/ozzo-validation/v4/is"
    "github.com/globalsign/mgo/bson"
    "time"
    "regexp"
)

type User struct {
    ID bson.ObjectId     `bson:"_id"`
    Name string          `bson:"name"`
    Email string         `bson:"email"`
    Password string      `bson:"password"`
    Token string         `bson:"token"`
    VerifiedAt time.Time `bson:"verified_at"`
    VerifyToken string   `bson:"verify_token"`
}

// func (user *User) CreateUserValidate() error {
//     email := &user.Email
//     return validation.Validate(email,
//         validation.Required.Error("メールアドレスは必須です"),
//         is.Email.Error("正しいメールアドレスの形で入力してください"),
//     )
// }

func (user *User) CreateUserValidate() error {
    return validation.ValidateStruct(user,
        validation.Field(&user.Email, validation.Required.Error("メールアドレスは必須です")),
        validation.Field(&user.Email, is.Email.Error("正しいメールアドレスの形で入力してください")),
        validation.Field(&user.Password, validation.Required.Error("パスワードは必須です")),
        validation.Field(&user.Password, validation.Match(regexp.MustCompile("^([a-zA-Z0-9]{8,})$")).Error("パスワードは半角英数8文字以上です")),
        validation.Field(&user.Password, validation.Length(8, 20).Error("パスワードは半角英数8文字以上です")),
    )
}