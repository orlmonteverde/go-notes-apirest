package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orlmonteverde/go-notes-apirest/models"
	"github.com/orlmonteverde/go-notes-apirest/models/notes"
)

// GetNotesHandler Handle Queries and response all notes like json
func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j, status := notes.GetAllNotes()
	w.WriteHeader(status)
	w.Write(j)
}

// GetNoteHandler Handle Queries and response one note like json
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	j, status := notes.GetNoteById(id)
	w.WriteHeader(status)
	w.Write(j)
}

// GetNoteHandler Handle Queries and response one note like json
func GetUserNotesHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	j, status := notes.GetNotesByUserId(id)
	w.WriteHeader(status)
	w.Write(j)
}

// PostNoteHandler Create new note and response a status code
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Println("Error al parsear nota")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	status := notes.PostNote(note)
	w.WriteHeader(status)
}

// DeleteNoteHandler Delete a note and response a status code
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	status := notes.DeleteNote(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

// PutNoteHandler Update a note and response a status code
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var noteUpdate models.Note
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if err != nil {
		log.Printf("Error al parsear nota con el id %s", id)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	status := notes.PutNote(id, noteUpdate)
	w.WriteHeader(status)
}
