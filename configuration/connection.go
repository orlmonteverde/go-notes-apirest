package configuration

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/orlmonteverde/go-notes-apirest/commons"
)

type configuration struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func getConfiguration() configuration {
	var c configuration
	file, err := os.Open("./configuration/config.json")
	commons.CheckError(err, "", true)
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	commons.CheckError(err, "", true)
	return c
}

// GetConnection return a connection to de database
func GetConnection() *sql.DB {
	c := getConfiguration()
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.Database,
	)
	db, err := sql.Open("postgres", dns)
	commons.CheckError(err, "", true)
	return db
}
