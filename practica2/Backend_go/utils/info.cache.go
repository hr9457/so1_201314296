package utils

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// funcion que retorna la memoria cache
func CacheMemory() int {

	cmd := exec.Command("sh", "-c", "free -m |head -n 2 |tail -1 |awk {'print $6'}")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	output := strings.Trim(string(out[:]), "\n")
	result, err := strconv.Atoi(output)

	if err != nil {
		fmt.Println("Error Memory Cache")
		return 0
	}
	return result
}
