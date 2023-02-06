package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hr9457/so1/routes"
)

func main() {

	port := ":4000"

	route := mux.NewRouter()

	// route principal
	routes.RoutesCalculadora(route)

	// Escucha del server
	log.Printf("Server listening on port %v", port)
	http.ListenAndServe(port, route)
}
