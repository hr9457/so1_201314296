package models

type Historial struct {
	Id_historial int    `json:"id_historial"`
	Numero1      int    `json:"numero1"`
	Numero2      int    `json:"numero2"`
	Operador     string `json:"operador"`
	Resultado    int    `json:"resultado"`
	Fecha        string `json:"fecha"`
}

type Cantidad struct {
	Id_historial int `json:"id_historial"`
	Cantidad     int `json:"cantidad"`
}

type CantidadError struct {
	Id_historial int    `json:"id_historial"`
	Numero1      int    `json:"numero1"`
	Numero2      int    `json:"numero2"`
	Operador     string `json:"operador"`
	Resultado    int    `json:"resultado"`
	Fecha        string `json:"fecha"`
}

type CantidadOperaciones struct {
	Operacion string `json:"operacion"`
	Cantidad  int    `json:"cantidad"`
}
