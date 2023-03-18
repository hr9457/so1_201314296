package utils

import (
	"api_go/database"
	"api_go/models"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// porcentaje cpu usado
func cpu() models.Usage {

	cmd := exec.Command("sh", "-c", "top -bn 1 -i -c | head -n 3 | tail -1 | awk {'print $8'}")
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	output := string(out[:])
	//eliminando espacios en blanco
	splitOutput := strings.Replace(output, "\n", "", -1)
	//eliminando saltoas de linea y convirtiendo en array

	var uso float64
	uso, err = strconv.ParseFloat(splitOutput, 64)

	if err != nil {
		fmt.Println("error de conversion", err)
		uso = 0
	}

	uso = 100 - uso
	var porcentaje models.Usage
	porcentaje.Porcentaje = uso

	return porcentaje
}

func returnUsername(uid int) string {
	cmd := exec.Command("sh", "-c", "getent passwd "+strconv.Itoa(uid)+" | cut -d: -f1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	output := string(out[:])
	output = strings.Replace(output, "\n", "", -1)
	return output
}

// funcion que evalua en que estado se encuentran los procesos
func estados(numEstado string) string {
	var tipoEstado string
	if numEstado == "0" {
		tipoEstado = "running"
	} else if numEstado == "1" || numEstado == "1026" || numEstado == "2" {
		tipoEstado = "sleeping"
	} else if numEstado == "4" {
		tipoEstado = "zombie"
	} else if numEstado == "8" {
		tipoEstado = "stopped"
	} else {
		tipoEstado = "no Reconocido"
	}
	return tipoEstado
}

// consumo del modulo de cpu
// (info models.InfoGeneral)
func InfoCpuProcess() {

	var procesoEjecucion int = 0
	var procesoSuspendido int = 0
	var procesoDetenido int = 0
	var procesoZombie int = 0

	cmd := exec.Command("sh", "-c", "cat /proc/cpu_201314296")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	output := string(out[:])

	//convirtiendo json a struct
	moduloCpu := []models.Data{}

	err2 := json.Unmarshal([]byte(output), &moduloCpu)
	if err2 != nil {
		fmt.Println("Error al leer json modulo de cpu", err2)
		//return moduloCpu
	}

	moduloCpu = moduloCpu[1:]
	//ciclo que formatea el json con el nombre del usuario

	for i, s := range moduloCpu {
		_ = i
		moduloCpu[i].Nombre = returnUsername(s.UID)
		moduloCpu[i].Estado = estados(moduloCpu[i].Estado)

		if moduloCpu[i].Estado == "running" {
			procesoEjecucion += 1
		} else if moduloCpu[i].Estado == "sleeping" {
			procesoSuspendido += 1
		} else if moduloCpu[i].Estado == "zombie" {
			procesoZombie += 1
		} else if moduloCpu[i].Estado == "stopped" {
			procesoDetenido += 1
		}

		moduloCpu[i].Hijos = moduloCpu[i].Hijos[1:]
		for j, hj := range moduloCpu[i].Hijos {
			_ = j
			nomH := returnUsername(hj.UID)
			moduloCpu[i].Hijos[j].Nombre = nomH
			moduloCpu[i].Hijos[j].Estado = estados(moduloCpu[i].Hijos[j].Estado)

			if moduloCpu[i].Hijos[j].Estado == "running" {
				procesoEjecucion += 1
			} else if moduloCpu[i].Hijos[j].Estado == "sleeping" {
				procesoSuspendido += 1
			} else if moduloCpu[i].Hijos[j].Estado == "zombie" {
				procesoZombie += 1
			} else if moduloCpu[i].Hijos[j].Estado == "stopped" {
				procesoDetenido += 1
			}
		}
	}

	total := procesoEjecucion + procesoSuspendido + procesoZombie + procesoDetenido
	infodata, err := json.Marshal(moduloCpu)
	porcentajeCpu := cpu().Porcentaje

	if err != nil {
		log.Fatal("Error convert to json ", err)
	}

	// insert database
	db, err := database.GetDB()
	if err != nil {
		log.Fatal("Error conecction database")
	}
	_, err = db.Exec("INSERT INTO cpu (porcentaje, ejecuccion, suspendidos, detenidos, zombies, total, data) VALUES(?,?,?,?,?,?,?)",
		porcentajeCpu, procesoEjecucion, procesoSuspendido, procesoZombie, procesoDetenido, total, infodata)

	if err != nil {
		log.Fatal("Error insert data cpu on database", err)
	}

	// return resultados
}
