package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orlmonteverde/gyga/apirest/models"
)

// GetUsers ...
func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j := models.GetNotesHandler()
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	j, found := models.GetNoteHandler(k)
	if !found {
		log.Printf("Id %s no encontrado", k)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// PostUsers ...
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Println("Error al recibir usuario")
	}
	w.Header().Set("Content-Type", "application/json")
	j := models.PostNoteHandler(note)
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// DeleteUser ...
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]

	if !models.DeleteNoteHandler(k) {
		fmt.Fprintf(w, "Id %s no encontrado", k)
	}

	w.WriteHeader(http.StatusNoContent)
}

// PutUser ...
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	var noteUpdate models.Note

	err := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if err != nil {
		log.Printf("Usuario con el id %s no encontrado", k)
	}

	if !models.PutNoteHandler(k, noteUpdate) {
		fmt.Fprintf(w, "Id %s no encontrado", k)
	}

	w.WriteHeader(http.StatusNoContent)
}
