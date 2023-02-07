package controller

import (
	"log"

	"github.com/hr9457/so1/database"
	"github.com/hr9457/so1/models"
)

// func de prueba
func HelloWorld() string {
	return "Hello World!!!"
}

// funciones de la calculadora
func Suma(op models.Operacion) models.ResponseOperacion {
	var response models.ResponseOperacion
	response.Result = op.Value1 + op.Value2

	// insert to database
	insertData("suma", response.Result)

	return response
}

func Resta(op models.Operacion) models.ResponseOperacion {
	var response models.ResponseOperacion
	response.Result = op.Value1 - op.Value2

	// insert to database
	insertData("resta", response.Result)

	return response
}

func Multiplicacion(op models.Operacion) models.ResponseOperacion {
	var response models.ResponseOperacion
	response.Result = op.Value1 * op.Value2

	// insert to database
	insertData("multiplicacion", response.Result)

	return response
}

func Division(op models.Operacion) models.ResponseOperacion {
	var response models.ResponseOperacion

	var value1 = op.Value1
	var value2 = op.Value2
	if value2 != 0 {
		response.Result = value1 / value2

		// insert to database
		insertData("division", response.Result)

		return response
	}

	// Err division
	response.Message = "Error division por 0"
	// insert to database
	insertData("Error", response.Result)
	return response
}

func insertData(tipo string, response int) {
	db, err := database.GetDB()
	if err != nil {
		log.Fatal("Error conecction database")
	}
	_, err = db.Exec("INSERT INTO history (tipo, resultado) VALUES(?,?)", tipo, response)
}
