package controller

import (
	"log"

	"github.com/hr9457/so1/database"
	"github.com/hr9457/so1/models"
	"github.com/hr9457/so1/util"
)

// func de prueba
func HelloWorld() string {
	return "Hello World!!!"
}

// funciones de la calculadora
func Suma(op models.Operacion) models.ResponseOperacion {
	var response models.ResponseOperacion
	response.Result = op.Value1 + op.Value2
	// fmt.Println(response)
	// insert to database
	insertData(op.Value1, op.Value2, "+", response.Result, util.GetDate())

	return response
}

func Resta(op models.Operacion) models.ResponseOperacion {
	var response models.ResponseOperacion
	response.Result = op.Value1 - op.Value2

	// insert to database
	insertData(op.Value1, op.Value2, "-", response.Result, util.GetDate())

	return response
}

func Multiplicacion(op models.Operacion) models.ResponseOperacion {
	var response models.ResponseOperacion
	response.Result = op.Value1 * op.Value2

	// insert to database
	insertData(op.Value1, op.Value2, "*", response.Result, util.GetDate())

	return response
}

func Division(op models.Operacion) models.ResponseOperacion {
	var response models.ResponseOperacion

	var value1 = op.Value1
	var value2 = op.Value2
	if value2 != 0 {
		response.Result = value1 / value2

		// insert to database
		insertData(value1, value2, "/", response.Result, util.GetDate())

		return response
	}

	// Err division
	response.Message = "Error division por 0"
	// insert to database
	insertData(value1, value2, "Error", response.Result, util.GetDate())
	return response
}

func insertData(num1 int, num2 int, op string, response int, fecha string) {
	db, err := database.GetDB()
	if err != nil {
		log.Fatal("Error conecction database")
	}
	_, err = db.Exec("INSERT INTO historial (numero1, numero2, operador, resultado, fecha) VALUES(?,?,?,?,?)", num1, num2, op, response, fecha)
}
