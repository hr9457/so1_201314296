package controller

import (
	"encoding/json"
	"net/http"

	"github.com/hr9457/so1/models"
)

func HelloWorld() string {
	return "Hello World!!!"
}

func Suma(w http.ResponseWriter, r *http.Request) string {
	var op models.Operacion
	json.NewDecoder(r.Body).Decode(&op)
	return "Hello Cal!!!"
}

func Resta() {}

func Multiplicacion() {}

func Division() {}
