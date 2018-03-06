package expense

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/go-playground/validator.v9"
	"github.com/fatih/structs"

	. "github.com/ziken/Register-of-Expenses-REST-API-go/db"
)

type Expense struct {
	Id      bson.ObjectId  `bson:"_id" json:"_id"`
	Title   string  `bson:"title,omitempty" json:"title" validate:"required,gt=2"`
	Amount  float32 `bson:"amount,omitempty" json:"amount" validate:"required,gt=0.01"`
	SpentAt int     `bson:"spentAt,omitempty" json:"spentAt" validate:"required"`
}

var validate * validator.Validate

func (ex *Expense) Validate() error {
	return validate.Struct(ex)
}

func (ex *Expense) ValidatePartial() error {

	fields := structs.New(ex).Fields()

	var validateFields []string
	for _, field := range fields {
		if !field.IsZero() {
			validateFields = append(validateFields, field.Name())
		}
	}

	if len(validateFields) == 0 {
		return errors.New("no fields to update")
	}
	return validate.StructPartial(ex, validateFields...)
}
//--------------------------//

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

func UpdateById(id string, expDoc Expense) error {
	expDoc.Id = bson.ObjectIdHex(id)
	return DB.C(EXPENSE_COLLECTION).UpdateId(expDoc.Id, bson.M{"$set": expDoc})
}

func init() {
	validate = validator.New()
}