package utils

import (
	"api_go/database"
	"api_go/models"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// models.MemoryRam

func MemoryModuleReturn() {
	var total int
	var consumida int
	var libre int
	var info models.MemoryRam

	// fmt.Println("ejecucion")
	cmd := exec.Command("sh", "-c", "cat /proc/ram_201314296")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	output := string(out[:])

	memory := models.ModuleRam{}

	err2 := json.Unmarshal([]byte(output), &memory)
	if err2 != nil {
		fmt.Println("Error al leer json modulo de memoria")
		// return info
	}

	total = memory.Total
	libre = memory.Free + CacheMemory()
	consumida = total - libre
	percent := (float64(consumida) / float64(total)) * 100

	info.TotalMemory = total
	info.ConsumedMemory = consumida
	info.UseMemory = percent

	// insert database
	db, err := database.GetDB()
	if err != nil {
		log.Fatal("Error conecction database")
	}
	_, err = db.Exec("INSERT INTO memory (total_memory, consumed_memory, use_memory) VALUES(?,?,?)", total, consumida, percent)
	if err != nil {
		log.Fatal("Error insert data memory on database")
	}

	// return info
}
