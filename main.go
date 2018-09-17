package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/orlmonteverde/go-apirest/helpers"
	"github.com/orlmonteverde/go-apirest/models"
	"github.com/orlmonteverde/go-apirest/routes"
)

var (
	defaultPort string
	server      *http.Server
	port        *string
	migrate     *bool
)

func main() {
	if *migrate {
		err := models.MakeMigrations()
		helpers.CheckError(err, "Fallaron las migraciones", true)
		log.Println("Migraciones realizadas con éxito")
	}
	log.Printf("Listening in port %s\n", *port)
	log.Fatal(server.ListenAndServe())

}

func init() {
	defaultPort = func() string {
		if envPort := os.Getenv("port"); envPort == "" {
			return "8080"
		} else {
			return envPort
		}
	}()

	port = flag.String("port", defaultPort, "Port to serve")
	migrate = flag.Bool("migrate", false, "Make Migrations")
	flag.Parse()

	r := routes.Mux

	server = &http.Server{
		Addr:           ":" + *port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
