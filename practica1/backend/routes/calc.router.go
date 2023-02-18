package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hr9457/so1/controller"
	"github.com/hr9457/so1/util"
)

// control routes calc
func RoutesCalculadora(route *mux.Router) {

	// route de prueba de saludo
	// route.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
	// 	var response = controller.HelloWorld()

	// 	// params url
	// 	fmt.Println(r.URL)
	// 	fmt.Println(r.URL.Query().Get("value1"))

	// 	w.Write([]byte(response))
	// }).Methods("POST")

	// route suma
	route.HandleFunc("/suma", func(w http.ResponseWriter, r *http.Request) {
		// response
		respondWhitSucces(controller.Suma(util.DestructuringData(*r)), w)
	}).Methods("POST")

	// Resta
	route.HandleFunc("/resta", func(w http.ResponseWriter, r *http.Request) {
		// response
		respondWhitSucces(controller.Resta(util.DestructuringData(*r)), w)
	}).Methods("POST")

	// Multiplicacion
	route.HandleFunc("/multiplicacion", func(w http.ResponseWriter, r *http.Request) {
		// response
		respondWhitSucces(controller.Multiplicacion(util.DestructuringData(*r)), w)
	}).Methods("POST")

	// Division
	route.HandleFunc("/division", func(w http.ResponseWriter, r *http.Request) {
		// response
		respondWhitSucces(controller.Division(util.DestructuringData(*r)), w)
	}).Methods("POST")

	// historial logs
	route.HandleFunc("/historial", func(w http.ResponseWriter, r *http.Request) {
		respondWhitSucces(controller.GetLogs(), w)
	}).Methods("GET")

	route.HandleFunc("/cantidad", func(w http.ResponseWriter, r *http.Request) {
		respondWhitSucces(controller.GetCantidad(), w)
	}).Methods("GET")

	route.HandleFunc("/cantidad/errores", func(w http.ResponseWriter, r *http.Request) {
		respondWhitSucces(controller.GetCantidadErr(), w)
	}).Methods("GET")

	route.HandleFunc("/cantidad/operaciones", func(w http.ResponseWriter, r *http.Request) {
		respondWhitSucces(controller.GetCantidadOp(), w)
	}).Methods("GET")

}

// func response succes api
func respondWhitSucces(data interface{}, w http.ResponseWriter) {
	// enable cors
	// fmt.Println(data)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
