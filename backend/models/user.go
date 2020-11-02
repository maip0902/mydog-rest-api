package models

import (
    "github.com/go-ozzo/ozzo-validation/v4"
    "github.com/go-ozzo/ozzo-validation/v4/is"
    "github.com/globalsign/mgo/bson"

)

type User struct {
    ID bson.ObjectId   `bson:"_id"`
    Name string        `bson:"name"`
    Email string       `bson:"email"`
    Password string    `bson:"password"`
}

func (user *User) CreateUserValidate() error {
    return validation.Validate(&user.Email,
        validation.Required.Error("メールアドレスは必須です"),
        is.Email.Error("正しいメールアドレスの形で入力してください"),
    )
}