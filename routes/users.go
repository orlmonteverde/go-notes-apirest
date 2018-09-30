package routes

import (
	"github.com/gorilla/mux"
	"github.com/orlmonteverde/go-notes-apirest/controllers"
	"github.com/orlmonteverde/go-notes-apirest/middlewares/auth"
)

func SetUsersRouter(router *mux.Router) {
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/",
		auth.ValidateAdmin(controllers.GetUsersHandler)).Methods("GET")
	subRouter.HandleFunc("/{id}",
		auth.ValidateUser(controllers.GetUserHandler)).Methods("GET")
	subRouter.HandleFunc("/", controllers.PostUserHandler).Methods("POST")
	subRouter.HandleFunc("/{id}",
		auth.ValidateUser(controllers.PutUserHandler)).Methods("PUT")
	subRouter.HandleFunc("/{id}", auth.ValidateUser(controllers.DeleteUserHandler)).Methods("DELETE")
	router.PathPrefix(prefix).Handler(subRouter)
}
