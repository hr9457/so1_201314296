package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hr9457/so1/controller"
	"github.com/hr9457/so1/util"
)

// control routes calc
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
		// response
		respondWhitSucces(controller.Suma(util.DestructuringData(*r)), w)
	}).Methods("GET")

	// Resta
	route.HandleFunc("/resta", func(w http.ResponseWriter, r *http.Request) {
		// response
		respondWhitSucces(controller.Resta(util.DestructuringData(*r)), w)
	}).Methods("GET")

	// Multiplicacion
	route.HandleFunc("/multiplicacion", func(w http.ResponseWriter, r *http.Request) {
		// response
		respondWhitSucces(controller.Multiplicacion(util.DestructuringData(*r)), w)
	}).Methods("GET")

	// Division
	route.HandleFunc("/division", func(w http.ResponseWriter, r *http.Request) {
		// response
		respondWhitSucces(controller.Division(util.DestructuringData(*r)), w)
	}).Methods("GET")

	// historial logs
	route.HandleFunc("/historial", func(w http.ResponseWriter, r *http.Request) {
		respondWhitSucces(controller.GetLogs(), w)
	}).Methods("GET")
}

// func enable control cors from frontend
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Acces-Control-Allow-Origin", "*")
}

// func response succes api
func respondWhitSucces(data interface{}, w http.ResponseWriter) {
	// enable cors
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
