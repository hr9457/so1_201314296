package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hr9457/so1/controller"
	"github.com/hr9457/so1/models"
)

func RoutesCalculadora(route *mux.Router) {

	// route de prueba de saludo
	route.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
		var response = controller.HelloWorld()

		// params url
		fmt.Println(r.URL)
		fmt.Println(r.URL.Query().Get("value1"))

		w.Write([]byte(response))
	}).Methods("GET")

	// route suma
	route.HandleFunc("/suma", func(w http.ResponseWriter, r *http.Request) {
		// model
		var op models.Operacion
		json.NewDecoder(r.Body).Decode(&op)

		// response
		response := controller.Suma(op)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	}).Methods("GET")

	// Resta
	route.HandleFunc("/resta", func(w http.ResponseWriter, r *http.Request) {
		// model
		var op models.Operacion
		json.NewDecoder(r.Body).Decode(&op)

		// response
		response := controller.Resta(op)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	}).Methods("GET")

	// Multiplicacion
	route.HandleFunc("/multiplicacion", func(w http.ResponseWriter, r *http.Request) {
		// model
		var op models.Operacion
		json.NewDecoder(r.Body).Decode(&op)

		// response
		response := controller.Multiplicacion(op)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	}).Methods("GET")

	// Division
	route.HandleFunc("/division", func(w http.ResponseWriter, r *http.Request) {
		// model
		var op models.Operacion
		json.NewDecoder(r.Body).Decode(&op)

		// response
		response := controller.Division(op)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	}).Methods("GET")

}
