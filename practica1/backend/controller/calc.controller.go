package controller

import (
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

	return response
}

func Resta(op models.Operacion) models.ResponseOperacion {
	var response models.ResponseOperacion
	response.Result = op.Value1 - op.Value2

	return response
}

func Multiplicacion(op models.Operacion) models.ResponseOperacion {
	var response models.ResponseOperacion
	response.Result = op.Value1 * op.Value2

	return response
}

func Division(op models.Operacion) models.ResponseOperacion {
	var response models.ResponseOperacion

	var value1 = op.Value1
	var value2 = op.Value2
	if value2 != 0 {
		response.Result = value1 / value2
		return response
	}

	response.Message = "Error division por 0"
	return response
}
