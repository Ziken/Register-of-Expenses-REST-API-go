package main
import (
	"net/http"
	"time"
	"log"
	"errors"
	"encoding/json"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"

	"github.com/ziken/Register-of-Expenses-REST-API-go/models/expense"
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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/expenses", nil).Methods("GET").HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		expenses, err := expense.FindAll()
		if checkErr(err, http.StatusBadRequest, w) {
			return
		}
		sendJSON(expenses, w)
	})

	r.HandleFunc("/expenses", nil).Methods("POST").HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
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
	})

	r.HandleFunc("/expenses/{id}", nil).Methods("GET").HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
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
	})

	r.HandleFunc("/expenses/{id}", nil).Methods("PATCH").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})

	r.HandleFunc("/expenses/{id}", nil).Methods("DELETE").HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
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
	})


	server := &http.Server{
		Addr: "127.0.0.1:3000",
		Handler: r,

		WriteTimeout: time.Second * 15,
		ReadTimeout: time.Second * 15,
		IdleTimeout: time.Second * 60,
	}
	c := make(chan int)

	go func(){
		log.Println("Server is running on port 3000");
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	<-c
}
