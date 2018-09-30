package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orlmonteverde/go-notes-apirest/commons"
	"github.com/orlmonteverde/go-notes-apirest/middlewares/auth"
	"github.com/orlmonteverde/go-notes-apirest/models"
	"github.com/orlmonteverde/go-notes-apirest/models/users"
)

// GetUsersHandler Handle Queries and response all user like json
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j, status := users.GetAllUsers()
	w.WriteHeader(status)
	w.Write(j)
}

// GetUserHandler Handle Queries and response one user like json
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	j, status := users.GetUserById(id)
	w.WriteHeader(status)
	w.Write(j)
}

// PostUserHandler Create new user and response a status code
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error al parsear usuario")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	status := users.CreateUser(user)
	w.WriteHeader(status)
}

// DeleteUserHandler Delete a user and response a status code
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	status := users.DeleteUser(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}

// PutUserHandler Update a user and response a status code
func PutUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var userUpdate models.User
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&userUpdate)
	if err != nil {
		log.Printf("Error al parsear usuario con el id %s", id)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	status := users.PutUser(id, userUpdate)
	w.WriteHeader(status)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error al leer el usuario: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error al leer el usuario")
		return
	}
	err = commons.PreparePassword(&user, false)
	if err != nil {
		log.Printf("Error al leer el usuario: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error al leer el usuario")
		return
	}
	storedUser, status := users.GetUserByName(user.Name)
	if status != http.StatusOK {
		log.Printf("Usuario %s no encontrado\n", user.Name)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Usuario %s no encontrado\n", user.Name)
		return
	}
	if user.Password == storedUser.Password {
		storedUser.Password = ""
		token := auth.GenerateJWT(storedUser)
		result := models.RespondeToken{token}
		j, err := json.Marshal(result)
		if err != nil {
			fmt.Fprintf(w, "Error al generar el json")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Usuario o clave no válidos")
	}
}
