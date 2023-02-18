package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hr9457/so1/models"
)

// func for destructuring
func DestructuringData(r http.Request) models.Operacion {
	var op models.Operacion
	json.NewDecoder(r.Body).Decode(&op)
	// fmt.Println(r.Body)
	return op
}

// func for date
func GetDate() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d% 02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return fecha
}
