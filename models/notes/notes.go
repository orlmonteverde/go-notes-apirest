package notes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/orlmonteverde/go-apirest/models"
)

// GetUsers Returns all users
func GetNotesHandler() (j []byte, status int) {
	var notes []Note
	q := `SELECT
		id, title, description, created_at, updated_at
		FROM notes`
	db := models.GetConnection()
	defer db.Close()

	rows, err := db.Query(q)
	if err != nil {
		log.Println("Error al consultar usuarios", err)
		status = http.StatusServiceUnavailable
		return
	}
	n := Note{}
	for rows.Next() {
		rows.Scan(&n.ID, &n.Title, &n.Description, &n.CreatedAt, &n.UpdatedAt)
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

// GetUsers Returns a user with id
func GetNoteHandler(idStr string) (j []byte, status int) {
	var n Note
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Id incorrecto", err)
		status = http.StatusBadRequest
		return
	}

	q := `SELECT
		id, title, description, created_at, updated_at
		FROM notes WHERE id=$1`
	db := models.GetConnection()
	defer db.Close()
	rows, err := db.Query(q, id)
	defer rows.Close()
	if err != nil {
		log.Println("Error al consultar usuario", err)
		status = http.StatusNotFound
		return
	}

	if rows.Next() {
		rows.Scan(&n.ID, &n.Title, &n.Description, &n.CreatedAt, &n.UpdatedAt)
	}

	j, err = json.Marshal(n)
	if err != nil {
		log.Println("Error al parsear usuario", err)
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusOK
	return
}

// PostUsers Create a new user
func PostNoteHandler(note Note) (status int) {
	q := `INSERT INTO
			notes(title, description, updated_at)
			VALUES($1, $2, now())`
	db := models.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare(q)
	if err != nil {
		log.Println("Error al preparar consulta", err)
		status = http.StatusInternalServerError
		return
	}
	defer stmt.Close()
	r, err := stmt.Exec(note.Title, note.Description)
	if err != nil {
		log.Println("Error al insertar usuario", err)
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

// DeleteUser Delete a user with id
func DeleteNoteHandler(idStr string) (status int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Id incorrecto", err)
		status = http.StatusBadRequest
		return
	}

	q := `DELETE FROM notes WHERE id=$1`
	db := models.GetConnection()
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

// PutUser Update auser with id
func PutNoteHandler(idStr string, note Note) (status int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Id incorrecto", err)
		status = http.StatusBadRequest
		return
	}
	q := `UPDATE notes SET
			title=$1, description=$2, updated_at=now()
			WHERE id=$3`
	db := models.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare(q)
	defer stmt.Close()
	r, err := stmt.Exec(note.Title, note.Description, id)
	if err != nil {
		log.Println("Error al consultar usuario", err)
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
