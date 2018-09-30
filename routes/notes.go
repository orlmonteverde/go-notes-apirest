package routes

import (
	"github.com/gorilla/mux"
	"github.com/orlmonteverde/go-notes-apirest/controllers"
	"github.com/orlmonteverde/go-notes-apirest/middlewares/auth"
)

func SetNotesRouter(router *mux.Router) {
	prefix := "/api/notes"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/",
		auth.ValidateUser(controllers.GetNotesHandler)).Methods("GET")
	subRouter.HandleFunc("/user/{id}", auth.ValidateUser(controllers.GetUserNotesHandler)).Methods("GET")
	subRouter.HandleFunc("/{id}", auth.ValidateUser(controllers.GetNoteHandler)).Methods("GET")
	subRouter.HandleFunc("/", auth.ValidateUser(controllers.PostNoteHandler)).Methods("POST")
	subRouter.HandleFunc("/{id}", auth.ValidateUser(controllers.PutNoteHandler)).Methods("PUT")
	subRouter.HandleFunc("/{id}", auth.ValidateUser(controllers.DeleteNoteHandler)).Methods("DELETE")
	router.PathPrefix(prefix).Handler(subRouter)
}
