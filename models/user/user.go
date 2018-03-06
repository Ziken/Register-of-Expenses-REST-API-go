package user

import (
	"gopkg.in/mgo.v2/bson"

)

type User struct {
	Id       bson.ObjectId `bson:"_id" json:"_id"`
	Email    string        `bson:"email" json:"email" validate:"required,email"`
	Password string        `bson:"password" json:"password" validate:"required,gt=5"`
	Tokens   []Token       `bson:"tokens" json:"-"`
}
type Token struct {
	Access string `bson:"access"  json:"-"`
	Token  string `bson:"token" json:"-"`
}

