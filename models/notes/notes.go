package notes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/orlmonteverde/go-notes-apirest/configuration"
	"github.com/orlmonteverde/go-notes-apirest/models"
)

func getNotes(q string, args ...interface{}) (j []byte, status int) {
	var notes []models.Note
	db := configuration.GetConnection()
	defer db.Close()

	rows, err := db.Query(q, args...)
	if err != nil {
		log.Println("Error al consultar usuarios", err)
		status = http.StatusServiceUnavailable
		return
	}
	n := models.Note{}
	for rows.Next() {
		rows.Scan(&n.ID, &n.UserID, &n.Title, &n.Description, &n.CreatedAt, &n.UpdatedAt)
		notes = append(notes, n)
	}
	j, err = json.Marshal(notes)
	if err != nil {
		log.Println("Error al parsear usuarios", err)
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusOK
	return
}

// GetAllNotes Returns all Notes
func GetAllNotes() (j []byte, status int) {
	q := `SELECT
		id, user_id, title, description, created_at, updated_at
		FROM notes`

	j, status = getNotes(q)
	return
}

func GetNoteById(idStr string) (j []byte, status int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Id incorrecto", err)
		status = http.StatusBadRequest
		return
	}

	q := `SELECT
		id, user_id, title, description, created_at, updated_at
		FROM notes WHERE id=$1`
	j, status = getNotes(q, id)
	return
}

func GetNotesByUserId(idStr string) (j []byte, status int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Id incorrecto", err)
		status = http.StatusBadRequest
		return
	}

	q := `SELECT
		id, user_id, title, description, created_at, updated_at
		FROM notes WHERE user_id=$1`
	j, status = getNotes(q, id)
	return
}

// PostNote Create a new note
func PostNote(n models.Note) (status int) {
	q := `INSERT INTO
			notes(user_id, title, description, updated_at)
			VALUES($1, $2, $3, now())`
	db := configuration.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare(q)
	if err != nil {
		log.Println("Error al preparar consulta", err)
		status = http.StatusInternalServerError
		return
	}
	defer stmt.Close()
	r, err := stmt.Exec(n.UserID, n.Title, n.Description)
	if err != nil {
		log.Println("Error al insertar nota", err)
		status = http.StatusBadRequest
		return
	}
	i, err := r.RowsAffected()
	if err != nil || i != 1 {
		log.Println("Se esperaba una fila afectada")
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusCreated
	return
}

// DeleteNote Delete a user for id
func DeleteNote(idStr string) (status int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Id incorrecto", err)
		status = http.StatusBadRequest
		return
	}

	q := `DELETE FROM notes WHERE id=$1`
	db := configuration.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare(q)
	defer stmt.Close()
	if err != nil {
		log.Println("Error al preparar consulta", err)
		status = http.StatusInternalServerError
		return
	}

	r, err := stmt.Exec(id)
	if err != nil {
		log.Println("Error al borrar nota", err)
		status = http.StatusNotFound
		return
	}

	i, err := r.RowsAffected()
	if err != nil || i != 1 {
		log.Println("Se esperaba una fila afectada")
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusNoContent
	return
}

// PutNote Update auser with id
func PutNote(idStr string, n models.Note) (status int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Id incorrecto", err)
		status = http.StatusBadRequest
		return
	}
	q := `UPDATE notes SET
			title=$1, description=$2, updated_at=now()
			WHERE id=$3`
	db := configuration.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare(q)
	defer stmt.Close()
	r, err := stmt.Exec(n.Title, n.Description, id)
	if err != nil {
		log.Println("Error al consultar nota", err)
		status = http.StatusInternalServerError
		return
	}
	i, err := r.RowsAffected()
	if err != nil || i == 0 {
		log.Println("Nota no existente")
		status = http.StatusBadRequest
		return
	} else if i > 1 {
		log.Println("Se esperaba una fila afectada")
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusNoContent
	return
}
