package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orlmonteverde/go-apirest/models/notes"
)

// GetNotesHandler Handle Queries and response all user like json
func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	var j []byte
	w.Header().Set("Content-Type", "application/json")
	j, status := notes.GetNotesHandler()
	w.WriteHeader(status)
	w.Write(j)
}

// GetNoteHandler Handle Queries and response one user like json
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	j, status := notes.GetNoteHandler(id)
	w.WriteHeader(status)
	w.Write(j)
}

// PostNoteHandler Create new user and response a status code
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note notes.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Println("Error al parsear usuario")
	}
	w.Header().Set("Content-Type", "application/json")
	status := notes.PostNoteHandler(note)
	w.WriteHeader(status)
}

// DeleteNoteHandler Delete a user and response a status code
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	status := notes.DeleteNoteHandler(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

// DeleteNoteHandler Update a user and response a status code
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var noteUpdate notes.Note

	err := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if err != nil {
		log.Printf("Error al parsear Usuario con el id %s", id)
	}
	w.Header().Set("Content-Type", "application/json")
	status := notes.PutNoteHandler(id, noteUpdate)
	w.WriteHeader(status)
}
