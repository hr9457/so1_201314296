package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hr9457/so1/routes"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", routes.Hello)

	http.ListenAndServe(":4000", r)
}
