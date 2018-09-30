package users

import (
	"log"
	"net/http"

	"github.com/orlmonteverde/go-notes-apirest/configuration"
	"github.com/orlmonteverde/go-notes-apirest/models"
)

func getUsers(q string, args ...interface{}) (users []models.User, err error) {
	db := configuration.GetConnection()
	defer db.Close()

	rows, err := db.Query(q, args...)
	if err != nil {
		log.Println("Error al consultar usuarios", err)
		return
	}
	u := models.User{}
	for rows.Next() {
		rows.Scan(&u.ID, &u.Name, &u.Role, &u.CreatedAt, &u.UpdatedAt)
		users = append(users, u)
	}
	return
}

func getUser(q string, args ...interface{}) (u models.User, err error) {
	db := configuration.GetConnection()
	defer db.Close()

	err = db.QueryRow(q, args...).Scan(&u.ID, &u.Name, &u.Password,
		&u.Role, &u.CreatedAt, &u.UpdatedAt)
	return
}

func postUser(q string, args ...interface{}) (status int) {

	db := configuration.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare(q)
	if err != nil {
		log.Println("Error al preparar consulta", err)
		status = http.StatusInternalServerError
		return
	}
	defer stmt.Close()

	r, err := stmt.Exec(args...)
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
