package models

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/orlmonteverde/go-apirest/helpers"
)

// GetConnection return a connection to de database
func GetConnection() *sql.DB {
	dns := "postgres://testadmin:testadmin@127.0.0.1:5432/gonotes?sslmode=disable"
	db, err := sql.Open("postgres", dns)
	helpers.CheckError(err, "", true)
	return db
}
