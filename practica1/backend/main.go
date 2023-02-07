package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hr9457/so1/database"
	"github.com/hr9457/so1/routes"
)

func main() {

	// port server
	port := ":4000"

	// route general
	route := mux.NewRouter()
	// Index routes
	prefix := route.PathPrefix("/api").Subrouter()
	// route principal
	routes.RoutesCalculadora(prefix)

	// prueba de db
	database.TestConnection()

	// Escucha del server
	log.Printf("Server listening on port %v", port)
	http.ListenAndServe(port, route)
}
