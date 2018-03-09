package authenticate

import (
	"net/http"
	"gopkg.in/mgo.v2/bson"
	. "github.com/ziken/Register-of-Expenses-REST-API-go/db"
	"github.com/ziken/Register-of-Expenses-REST-API-go/models/user"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-auth")
		var usr user.User
		if token == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)
			return
		}
		err := DB.C(USER_COLLECTION).Find(bson.M{
			"tokens.token": token,
		}).One(&usr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(nil)
			return
		}
		r.Header.Set("x-s-user-id", usr.Id.Hex())
		r.Header.Set("x-s-user-email", usr.Email)

		next.ServeHTTP(w, r)
	})
}
