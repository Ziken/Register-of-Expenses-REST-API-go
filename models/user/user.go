package user

import (
	"gopkg.in/mgo.v2/bson"

	. "github.com/ziken/Register-of-Expenses-REST-API-go/db"

	"gopkg.in/go-playground/validator.v9"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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


var validate * validator.Validate

func (usr * User) Validate() error {
	return validate.Struct(usr)
}

func (usr *  User) GenerateAuthToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": usr.Id,
		"access": "auth",
	})

	tokenString, err := token.SignedString([]byte("good-secret"))
	if err != nil {
		return "", err
	}

	err = DB.C(USER_COLLECTION).UpdateId(usr.Id, bson.M{
		"$push": bson.M{
			"tokens": Token{Token:tokenString, Access:"auth"},
		},
	})

	return tokenString, err
}
func (usr *  User) RemoveAuthToken(token string) (error) {
	return DB.C(USER_COLLECTION).UpdateId(usr.Id, bson.M{
		"$pull": bson.M{
			"tokens": bson.M{
				"token": token,
			},
		},
	})
}

func Save(userDoc User) (User, error) {
	userDoc.Id = bson.NewObjectId()
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userDoc.Password), bcrypt.DefaultCost)
	if err != nil {
		return userDoc, err
	}
	userDoc.Password = string(hashedPass)

	err = DB.C(USER_COLLECTION).Insert(userDoc)
	return userDoc, err;
}
func FindByToken(token string) (User,  error) {
	var usr User
	err := DB.C(USER_COLLECTION).Find(bson.M{
		"tokens.token": token,
	}).One(&usr)

	return usr, err
}
func FindByCredentials(email, password string) (User, error) {
	var usr User
	err := DB.C(USER_COLLECTION).Find(bson.M{
		"email": email,
	}).One(&usr)

	if err != nil {
		return usr, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	return usr, err
}
func init() {
	validate = validator.New()
}
