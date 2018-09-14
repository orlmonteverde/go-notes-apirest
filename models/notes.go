package models

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

var Id int

// GetUsers ...
func GetNotesHandler() []byte {
	var notes []Note
	for _, v := range NoteStore {
		notes = append(notes, v)
	}
	j, err := json.Marshal(notes)
	if err != nil {
		log.Println("Error al leer usuarios")
	}
	return j
}

func GetNoteHandler(k string) (j []byte, found bool) {
	var note Note
	note, found = NoteStore[k]

	j, err := json.Marshal(note)
	if err != nil || !found {
		log.Println("Usuario no encontrado")
	}
	return
}

// PostUsers ...
func PostNoteHandler(note Note) []byte {
	note.CreatedAt = time.Now()
	Id++
	k := strconv.Itoa(Id)
	NoteStore[k] = note
	j, err := json.Marshal(note)
	if err != nil {
		log.Fatal("Error al recibir usuario")
	}
	return j
}

// DeleteUser ...
func DeleteNoteHandler(id string) (found bool) {
	if _, ok := NoteStore[id]; ok {
		found = true
		delete(NoteStore, id)
	}
	return
}

// PutUser ...
func PutNoteHandler(id string, noteUpdate Note) (found bool) {
	if note, ok := NoteStore[id]; ok {
		found = true
		noteUpdate.CreatedAt = note.CreatedAt
		NoteStore[id] = noteUpdate
	}
	return
}
