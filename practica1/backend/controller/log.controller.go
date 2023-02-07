package controller

import (
	"github.com/hr9457/so1/database"
	"github.com/hr9457/so1/models"
)

func GetLogs() []models.Logs {
	historial := []models.Logs{}

	db, err := database.GetDB()
	if err != nil {
		return historial
	}

	rows, err := db.Query("SELECT tipo, resultado FROM history")
	if err != nil {
		return historial
	}

	// Iteracion
	for rows.Next() {
		var logs models.Logs
		err = rows.Scan(&logs.Tipo, &logs.Resultado)
		if err != nil {
			return historial
		}
		historial = append(historial, logs)
	}

	return historial

}
