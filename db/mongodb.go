package db

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

type DbHandler struct {
	Server   string
	Database string
}

var DB *mgo.Database
const EXPENSE_COLLECTION = "expenses"
const USER_COLLECTION = "users"

func (e *DbHandler) Connect() {
	session, err := mgo.Dial(e.Server)
	if err != nil {
		log.Fatal(err)
	}
	DB = session.DB(e.Database)
}

func init () { // TODO load config
	var dbHandler DbHandler
	dbHandler.Server = os.Getenv("MONGOHQ_SERVER")
	dbHandler.Database = os.Getenv("MONGOHQ_DATABASE")
	//dbHandler.Server = "127.0.0.1"
	//dbHandler.Database = "ExpensesRegister"
	dbHandler.Connect()
}
