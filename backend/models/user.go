package models

import (
    "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/go-ozzo/ozzo-validation/v4/is"
    "github.com/globalsign/mgo/bson"
    "time"
)

type User struct {
    ID bson.ObjectId     `bson:"_id"`
    Name string          `bson:"name"`
    Email string         `bson:"email"`
    Password string      `bson:"password"`
    Token string         `bson:"token"`
    VerifiedAt time.Time `bson:"verified_at"`
}

func (user *User) CreateUserValidate() error {
    email := &user.Email
    return validation.Validate(email,
        validation.Required.Error("メールアドレスは必須です"),
        is.Email.Error("正しいメールアドレスの形で入力してください"),
    )
}