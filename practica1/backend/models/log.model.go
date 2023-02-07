package models

type Historial struct {
	Numero1   int    `json:"numero1"`
	Numero2   int    `json:"numero2"`
	Operador  string `json:"operador"`
	Resultado int    `json:"resultado"`
	Fecha     string `json:"fecha"`
}
