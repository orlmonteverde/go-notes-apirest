package models

import (
	"io/ioutil"

	"github.com/orlmonteverde/go-apirest/helpers"
)

// MakeMigrations Create all tables at database
func MakeMigrations() error {
	file, err := ioutil.ReadFile("./models/models.sql")
	helpers.CheckError(err, "Archivo no encontrado", true)
	q := string(file)

	db := GetConnection()
	defer db.Close()
	rows, err := db.Query(q)
	defer rows.Close()
	helpers.CheckError(err, "Error al realizar la consulta", true)
	return nil
}
