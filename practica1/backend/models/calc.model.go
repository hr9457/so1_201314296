package models

// Opercion -> struct para el manejo de valores para la operaciones
type Operacion struct {
	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
}

type ResponseOperacion struct {
	Result  int    `json:"result"`
	Message string `json:"message"`
}
