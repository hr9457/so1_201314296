package routes

import (
	"api_go/controller"
	"api_go/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RoutesApi(route *mux.Router) {

	route.HandleFunc("/saludo", func(w http.ResponseWriter, r *http.Request) {
		var response = controller.GetMemory()
		w.Write([]byte(response))
	}).Methods("GET")

	route.HandleFunc("/memory", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("llamando al moudlo ram")
		responseWhitSucess(utils.MemoryModuleReturn(), w)
	}).Methods("GET")

	route.HandleFunc("/cpu", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("llamando al moudlo cpu")
		responseWhitSucess(utils.InfoCpuProcess(), w)
	}).Methods("GET")
}

func responseWhitSucess(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
