package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orlmonteverde/go-apirest/controllers/notes"
)

var (
	Mux *mux.Router
)

func init() {
	fs := http.FileServer(http.Dir("public"))
	Mux = mux.NewRouter()
	Mux = mux.NewRouter().StrictSlash(false)
	Mux.Handle("/", fs)
	Mux.HandleFunc("/api/notes", controllers.GetNotesHandler).Methods("GET")
	Mux.HandleFunc("/api/notes/{id}", controllers.GetNoteHandler).Methods("GET")
	Mux.HandleFunc("/api/notes", controllers.PostNoteHandler).Methods("POST")
	Mux.HandleFunc("/api/notes/{id}", controllers.PutNoteHandler).Methods("PUT")
	Mux.HandleFunc("/api/notes/{id}", controllers.DeleteNoteHandler).Methods("DELETE")
}
