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
	Creator bson.ObjectId `bson:"_creator" json:"_creator"`
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

func FindAll(creator bson.ObjectId) ([]Expense, error) {
	var expenses []Expense
	err := DB.C(EXPENSE_COLLECTION).Find(bson.M{
		"_creator": creator,
	}).All(&expenses)
	return expenses, err
}

func FindById(id string, creator bson.ObjectId) (Expense, error) {
	var expDoc Expense
	err := DB.C(EXPENSE_COLLECTION).Find(bson.M{
		"_id": bson.ObjectIdHex(id),
		"_creator": creator,
	}).One(&expDoc)
	return expDoc, err
}

func RemoveById(id string, creator bson.ObjectId) (error) {
	err := DB.C(EXPENSE_COLLECTION).Remove(bson.M{
		"_id": bson.ObjectIdHex(id),
		"_creator": creator,
	})
	return err
}

func Save(expDoc Expense) (Expense, error) {
	expDoc.Id = bson.NewObjectId()
	err := DB.C(EXPENSE_COLLECTION).Insert(expDoc)
	return expDoc, err
}

func UpdateById(id string, expDoc Expense, creator bson.ObjectId) error {
	expDoc.Id = bson.ObjectIdHex(id)
	expDoc.Creator = creator
	return DB.C(EXPENSE_COLLECTION).Update(bson.M{
		"_id": expDoc.Id,
		"_creator": creator,
	}, bson.M{"$set": expDoc})
}

func init() {
	validate = validator.New()
}