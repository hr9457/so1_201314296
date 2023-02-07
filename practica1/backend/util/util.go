package util

import (
	"encoding/json"
	"net/http"

	"github.com/hr9457/so1/models"
)

// func for destructuring
func DestructuringData(r http.Request) models.Operacion {
	var op models.Operacion
	json.NewDecoder(r.Body).Decode(&op)
	return op
}
