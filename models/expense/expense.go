package expense

import (
	. "github.com/ziken/Register-of-Expenses-REST-API-go/db"
	"gopkg.in/mgo.v2/bson"
)

type Expense struct {
	Id      bson.ObjectId  `bson:"_id" json:"_id"`
	Title   string  `bson:"title,omitempty" json:"title" validate:"required,gt=2"`
	Amount  float32 `bson:"amount,omitempty" json:"amount" validate:"required,gt=0.01"`
	SpentAt int     `bson:"spentAt,omitempty" json:"spentAt" validate:"required"`
}

func FindAll() ([]Expense, error) {
	var expenses []Expense
	err := DB.C(EXPENSE_COLLECTION).Find(bson.M{}).All(&expenses)
	return expenses, err
}

func FindById(id string) (Expense, error) {
	var expDoc Expense
	err := DB.C(EXPENSE_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&expDoc)
	return expDoc, err
}

func RemoveById(id string) (error) {
	err := DB.C(EXPENSE_COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func Save(expDoc Expense) (Expense, error) {
	expDoc.Id = bson.NewObjectId()
	err := DB.C(EXPENSE_COLLECTION).Insert(expDoc)
	return expDoc, err
}