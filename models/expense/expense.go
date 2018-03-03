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
