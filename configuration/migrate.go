package configuration

import (
	"io/ioutil"

	"github.com/orlmonteverde/go-notes-apirest/commons"
)

// MakeMigrations Create all tables at database
func MakeMigrations() error {
	file, err := ioutil.ReadFile("./configuration/models.sql")
	commons.CheckError(err, "Archivo no encontrado", true)
	q := string(file)

	db := GetConnection()
	defer db.Close()
	rows, err := db.Query(q)
	defer rows.Close()
	commons.CheckError(err, "Error al realizar la consulta", true)
	return nil
}
