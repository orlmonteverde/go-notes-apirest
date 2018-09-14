package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/orlmonteverde/gyga/apirest/routes"
)

var (
	defaultPort string
	server      *http.Server
	port        *string
)

func main() {
	log.Printf("Listening in port %s\n", *port)
	log.Fatal(server.ListenAndServe())

}

func init() {
	defaultPort = func() string {
		if envPort := os.Getenv("port"); envPort == "" {
			return "8000"
		} else {
			return envPort
		}
	}()

	port = flag.String("port", defaultPort, "Port to serve")
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
