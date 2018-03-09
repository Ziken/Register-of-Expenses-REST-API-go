package authenticate

import (
	"net/http"
	"github.com/ziken/Register-of-Expenses-REST-API-go/models/user"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-auth")
		if token == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)
			return
		}
		usr, err := user.FindByToken(token)
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
