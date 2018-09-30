package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/orlmonteverde/go-notes-apirest/commons"
	"github.com/orlmonteverde/go-notes-apirest/configuration"
	"github.com/orlmonteverde/go-notes-apirest/routes"
)

var port string

func main() {

	migrate := flag.Bool("migrate", false, "Make Migrations")
	flag.Parse()

	if *migrate {
		err := configuration.MakeMigrations()
		commons.CheckError(err, "Fallaron las migraciones", true)
		log.Println("Migraciones realizadas con éxito")
	}

	if envPort := os.Getenv("port"); envPort == "" {
		port = "8080"
	} else {
		port = envPort
	}

	r := routes.InitRoutes()

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Listening in port %s\n", port)
	log.Fatal(server.ListenAndServe())

}
