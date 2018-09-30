package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orlmonteverde/go-notes-apirest/controllers"
)

// InitRoutes initialize all routes
func InitRoutes() *mux.Router {
	fs := http.FileServer(http.Dir("public"))

	router := mux.NewRouter().StrictSlash(false)
	SetNotesRouter(router)
	SetUsersRouter(router)
	router.HandleFunc("/login", controllers.Login)
	router.Handle("/", fs)
	return router
}
