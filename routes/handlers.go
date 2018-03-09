package routes

import (
	"errors"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"

	"github.com/ziken/Register-of-Expenses-REST-API-go/models/expense"
	"github.com/ziken/Register-of-Expenses-REST-API-go/models/user"
)
type ResponseJSON struct {
	Result interface{} `json:"result"`
}

func checkErr(err error, status int, w http.ResponseWriter) bool {
	if err != nil {
		w.WriteHeader(status)
		sendJSON(nil, w)
		return true
	}

	return false
}

func sendJSON(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type","application/json")

	if data == nil {
		w.Write(nil)
		return
	}
	structJSON := ResponseJSON{data}
	if err := json.NewEncoder(w).Encode(structJSON); err != nil {
		checkErr(err, http.StatusBadRequest, w)
	}
}

func GetExpenses(w http.ResponseWriter, r * http.Request) {
	expenses, err := expense.FindAll()
	if checkErr(err, http.StatusBadRequest, w) {
		return
	}
	sendJSON(expenses, w)
}

//r.HandleFunc("/expenses", nil).Methods("POST").HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
func PostExpense(w http.ResponseWriter, r * http.Request) {
	var expDoc expense.Expense

	err := json.NewDecoder(r.Body).Decode(&expDoc)
	if  checkErr(err, http.StatusBadRequest, w) {
		return
	}

	if  err := expDoc.Validate(); checkErr(err, http.StatusBadRequest, w) {
		return
	}

	insertedDoc, err := expense.Save(expDoc)
	if  checkErr(err, http.StatusBadRequest, w) {
		return
	}

	sendJSON(insertedDoc, w)
}

func GetExpenseById (w http.ResponseWriter, r * http.Request) {
	idExp := mux.Vars(r)["id"]
	if !bson.IsObjectIdHex(idExp) {
		checkErr(errors.New("invalid id"), http.StatusBadRequest, w)
		return
	}
	expDoc, err := expense.FindById(idExp)

	if err == mgo.ErrNotFound {
		checkErr(errors.New("not found"), http.StatusNotFound, w)
		return
	}
	if checkErr(err, http.StatusBadRequest, w) {
		return
	}

	sendJSON(expDoc, w)
}

func PatchExpenseById(w http.ResponseWriter, r *http.Request) {
	idExp := mux.Vars(r)["id"]
	if !bson.IsObjectIdHex(idExp) {
		checkErr(errors.New("invalid id"), http.StatusBadRequest, w)
		return
	}

	var expDoc expense.Expense
	if  err := json.NewDecoder(r.Body).Decode(&expDoc); checkErr(err, http.StatusBadRequest, w) {
		return
	}
	if err := expDoc.ValidatePartial(); checkErr(err, http.StatusBadRequest, w) {
		return
	}
	if err := expense.UpdateById(idExp, expDoc); checkErr(err, http.StatusBadRequest, w) {
		return
	}
	sendJSON(nil, w)
}

func DeleteExpenseById(w http.ResponseWriter, r * http.Request) {
	idExp := mux.Vars(r)["id"]
	if !bson.IsObjectIdHex(idExp) {
		checkErr(errors.New("invalid id"), http.StatusBadRequest, w)
		return
	}

	err := expense.RemoveById(idExp);
	if err == mgo.ErrNotFound {
		checkErr(errors.New("not found"), http.StatusNotFound, w)
		return
	}
	if checkErr(err, http.StatusBadRequest, w) {
		return ;
	}
	sendJSON(nil, w)
}

func PostUser(w http.ResponseWriter, r * http.Request) {
	var userDoc user.User

	err := json.NewDecoder(r.Body).Decode(&userDoc)

	if  checkErr(err, http.StatusBadRequest, w) {
		return
	}
	if err := userDoc.Validate(); checkErr(err, http.StatusBadRequest, w) {
		return
	}

	insertedDoc, err := user.Save(userDoc);
	if  checkErr(err, http.StatusBadRequest, w) {
		//log.Println(err)
		return
	}
	token, err := insertedDoc.GenerateAuthToken()

	if  checkErr(err, http.StatusBadRequest, w) {
		//log.Println(err)
		return
	}
	//log.Println("TOKEN", token)

	w.Header().Set("x-auth", token)
	sendJSON(insertedDoc, w);
}

func GetUserMe(w http.ResponseWriter, r * http.Request) {
	var usr user.User

	usr.Email = r.Header.Get("x-s-user-email")
	usr.Id = bson.ObjectIdHex(r.Header.Get("x-s-user-id"))

	sendJSON(usr, w)
}