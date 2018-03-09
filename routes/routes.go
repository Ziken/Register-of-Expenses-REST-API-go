package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ziken/Register-of-Expenses-REST-API-go/middleware/authenticate"
)

type Route struct {
	Path string
	Method string
	Authenticate bool
	HandlerFunc http.HandlerFunc
}

type Routes []Route


var routes = Routes{
	Route{
		Path: "/expenses",
		Method: "GET",
		Authenticate: false,
		HandlerFunc: GetExpenses,
	},
	Route{
		Path: "/expenses",
		Method: "POST",
		Authenticate: false,
		HandlerFunc: PostExpense,
	},
	Route{
		Path: "/expenses/{id}",
		Method: "GET",
		Authenticate: false,
		HandlerFunc: GetExpenseById,
	},
	Route{
		Path: "/expenses/{id}",
		Method: "PATCH",
		Authenticate: false,
		HandlerFunc: PatchExpenseById,
	},

	Route{
		Path: "/expenses/{id}",
		Method: "DELETE",
		Authenticate: false,
		HandlerFunc: DeleteExpenseById,
	},
	// USERS
	Route{
		Path: "/users",
		Method: "POST",
		Authenticate: false,
		HandlerFunc: PostUser,
	},
	Route{
		Path: "/users/me",
		Method: "GET",
		Authenticate: true,
		HandlerFunc: GetUserMe,
	},
	Route{
		Path: "/users/login",
		Method: "POST",
		Authenticate: false,
		HandlerFunc: PostUserLogin,
	},
	Route{
		Path: "/users/logout",
		Method: "POST",
		Authenticate: true,
		HandlerFunc: GetUserLogout,
	},
}

func NewRouter() (* mux.Router) {
	mainRouter := mux.NewRouter()
	authRouter := mux.NewRouter()

	authRouter.Use(authenticate.Authenticate)

	for _, route := range routes {
		if route.Authenticate {
			authRouter.HandleFunc(route.Path, route.HandlerFunc).Methods(route.Method)
			mainRouter.Handle(route.Path, authRouter)
		} else {
			mainRouter.HandleFunc(route.Path, route.HandlerFunc).Methods(route.Method)
		}
	}
	return mainRouter
}



