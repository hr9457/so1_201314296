package controller

import (
	"github.com/hr9457/so1/database"
	"github.com/hr9457/so1/models"
)

func GetLogs() []models.Historial {
	historial := []models.Historial{}

	db, err := database.GetDB()
	if err != nil {
		return historial
	}

	rows, err := db.Query("SELECT numero1, numero2, operador, resultado, fecha FROM historial")
	if err != nil {
		return historial
	}

	// Iteracion
	for rows.Next() {
		var log models.Historial
		err = rows.Scan(&log.Numero1, &log.Numero2, &log.Operador, &log.Resultado, &log.Fecha)
		if err != nil {
			return historial
		}
		historial = append(historial, log)
	}

	return historial

}
