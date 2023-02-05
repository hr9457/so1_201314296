package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hr9457/so1/controller"
)

func Hello(route *mux.Router) {

	// route de prueba de saludo
	route.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
		var response = controller.HelloWorld()
		w.Write([]byte(response))
	}).Methods("GET")

}
