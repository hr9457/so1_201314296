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

	rows, err := db.Query("SELECT id_historial, numero1, numero2, operador, resultado, fecha FROM historial")
	if err != nil {
		return historial
	}

	// Iteracion
	for rows.Next() {
		var log models.Historial
		err = rows.Scan(&log.Id_historial, &log.Numero1, &log.Numero2, &log.Operador, &log.Resultado, &log.Fecha)
		if err != nil {
			return historial
		}
		historial = append(historial, log)
	}

	return historial

}

func GetCantidad() []models.Cantidad {
	historial := []models.Cantidad{}

	db, err := database.GetDB()
	if err != nil {
		return historial
	}

	rows, err := db.Query("SELECT COUNT(*) FROM historial;")
	if err != nil {
		return historial
	}

	// Iteracion
	for rows.Next() {
		var log models.Cantidad
		err = rows.Scan(&log.Cantidad)
		if err != nil {
			return historial
		}
		historial = append(historial, log)
	}

	return historial
}

func GetCantidadErr() []models.CantidadError {
	historial := []models.CantidadError{}

	db, err := database.GetDB()
	if err != nil {
		return historial
	}

	rows, err := db.Query("SELECT id_historial, numero1, numero2, operador, resultado, fecha FROM historial WHERE operador = \"Error\";")
	if err != nil {
		return historial
	}

	// Iteracion
	for rows.Next() {
		var log models.CantidadError
		err = rows.Scan(&log.Id_historial, &log.Numero1, &log.Numero2, &log.Operador, &log.Resultado, &log.Fecha)
		if err != nil {
			return historial
		}
		historial = append(historial, log)
	}

	return historial
}

func GetCantidadOp() []models.CantidadOperaciones {
	historial := []models.CantidadOperaciones{}

	db, err := database.GetDB()
	if err != nil {
		return historial
	}

	queryMultiple := "SELECT 'sumas' as operacion, COUNT(*) as cantidad FROM historial WHERE operador = \"+\" \n" +
		"union	SELECT 'restas' as operacion, COUNT(*) as cantidad FROM historial WHERE operador = \"-\" \n" +
		"union SELECT 'multiplicaciones' as operacion, COUNT(*) as cantidad FROM historial WHERE operador = \"*\" " +
		"union 	SELECT 'divisiones' as operacion, COUNT(*) as cantidad FROM historial WHERE operador = \"/\" ;"

	rows, err := db.Query(queryMultiple)
	if err != nil {
		return historial
	}

	// Iteracion
	for rows.Next() {
		var log models.CantidadOperaciones
		err = rows.Scan(&log.Operacion, &log.Cantidad)
		if err != nil {
			return historial
		}
		historial = append(historial, log)
	}

	return historial
}
