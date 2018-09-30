package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/orlmonteverde/go-notes-apirest/commons"
	"github.com/orlmonteverde/go-notes-apirest/configuration"
	"github.com/orlmonteverde/go-notes-apirest/models"
)

// GetAllUsers Returns an Slice with all users
func GetAllUsers() (j []byte, status int) {
	q := `SELECT
		id, name, role, created_at, updated_at
		FROM users`
	users, err := getUsers(q)
	if err != nil {
		log.Printf("Error al consultar usuario: %s", err)
		status = http.StatusNotFound
	}
	j, err = commons.JsonParser(users)
	if err != nil {
		log.Printf("Error al parsear usuario: %s", err)
		status = http.StatusInternalServerError
	}
	status = http.StatusOK
	return
}

// GetUserById Returns a user for id
func GetUserById(idStr string) (j []byte, status int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Id incorrecto", err)
		status = http.StatusBadRequest
		return
	}

	q := `SELECT
		id, name, password, role, created_at, updated_at
		FROM users WHERE id=$1`
	user, err := getUser(q, id)
	if err != nil {
		log.Printf("Error al consultar usuario: %s", err)
		status = http.StatusNotFound
	}
	user.Password = ""
	j, err = commons.JsonParser(user)
	if err != nil {
		log.Printf("Error al parsear usuario: %s", err)
		status = http.StatusInternalServerError
	}
	status = http.StatusOK
	return
}

// GetUserById Returns a user for id
func GetUserByName(name string) (u models.User, status int) {
	q := `SELECT
		id, name, password, role, created_at, updated_at
		FROM users WHERE name=$1`

	u, err := getUser(q, name)
	if err != nil {
		log.Printf("Error al consultar usuario: %s", err)
		status = http.StatusNotFound
	}
	status = http.StatusOK
	return
}

// GetUserById Returns a user for id
func GetUserByRole(role string) (j []byte, status int) {
	q := `SELECT
		id, name, password, role, created_at, updated_at
		FROM users WHERE role=$1`
	users, err := getUsers(q, role)
	if err != nil {
		log.Printf("Error al consultar usuario: %s", err)
		status = http.StatusNotFound
	}
	for _, user := range users {
		user.Password = ""
	}
	j, err = commons.JsonParser(users)
	if err != nil {
		log.Printf("Error al parsear usuario: %s", err)
		status = http.StatusInternalServerError
	}
	status = http.StatusOK
	return
}

func CreateUser(u models.User) (status int) {
	q := `INSERT INTO
			users(name, password, updated_at)
			VALUES($1, $2, now())`
	err := commons.PreparePassword(&u, true)
	if err != nil {
		log.Println("Error:", err)
		status = http.StatusBadRequest
		return
	}
	status = postUser(q, u.Name, u.Password)
	return
}

func CreateAdmin(u models.User) (status int) {
	q := `INSERT INTO
			users(name, password, role, updated_at)
			VALUES($1, $2, true, now())`
	err := commons.PreparePassword(&u, true)
	if err != nil {
		log.Println("Error las contraseñas no coinciden")
		status = http.StatusBadRequest
		return
	}
	status = postUser(q, u.Name, u.Password)
	return
}

// DeleteUser Delete a user for id
func DeleteUser(idStr string) (status int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Id incorrecto", err)
		status = http.StatusBadRequest
		return
	}

	q := `DELETE FROM users WHERE id=$1`
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
		log.Printf("Error al borrar usuario %d: %s\n", id, err.Error())
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

// PutUserHandler Update a user for id
func PutUser(idStr string, u models.User) (status int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Id incorrecto", err)
		status = http.StatusBadRequest
		return
	}

	q := `UPDATE users SET
			name=$1, password=$2, updated_at=now()
			WHERE id=$3`
	db := configuration.GetConnection()
	defer db.Close()
	stmt, err := db.Prepare(q)
	defer stmt.Close()

	commons.PreparePassword(&u, true)
	r, err := stmt.Exec(u.Name, u.Password, id)
	if err != nil {
		log.Println("Error al actualizar usuario", err)
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
