package server

import (
	"log"
	"net/http"
	"os"
)

func Start() {
	router := SetUpRoutes()
	port := os.Getenv("API_PORT")
	if port == "" {
		port = ":8080"
	}

	log.Println("El server está corriendo en el port: ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
