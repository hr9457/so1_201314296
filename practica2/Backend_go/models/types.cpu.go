package models

// utilizacion del procesador
type Usage struct {
	Porcentaje float64 `json:"porcentaje"`
}

// informacion del modulo de cpu
type Data struct {
	Nombre  string `json:"nombre"`
	UID     int    `json:"UID"`
	PID     int    `json:"PID"`
	Proceso string `json:"proceso"`
	Estado  string `json:"estado"`
	VmRSS   int    `json:"vmrss"`
	Hijos   []Data `json:"hijos"`
}

// get info modulo cpu
type InfoGeneral struct {
	ProcesoEjecucion  int    `json:"ejecucion"`
	ProcesoSuspendido int    `json:"suspendidos"`
	ProcesosDetenidos int    `json:"detenidos"`
	ProcesosZombies   int    `json:"zombies"`
	Total             int    `json:"total"`
	DataProcesos      []Data `json:"procesos"`
}
