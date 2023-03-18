package main

import (
	"api_go/database"
	"api_go/utils"
	"fmt"
	"time"
)

func main() {
	// port server
	// port := ":4000"

	// // router general
	// route := mux.NewRouter()
	// middleware.EnableCORS(route)

	// // route principal
	// routes.RoutesApi(route)

	// prueba de db
	database.TestConnection()

	// // Escucha del server
	// log.Printf("Server listening on port %v", port)
	// http.ListenAndServe(port, route)

	for {

		utils.MemoryModuleReturn()
		fmt.Println("insert memory")

		utils.InfoCpuProcess()
		fmt.Println("insert cpu")
		// sleep process for 10 second
		time.Sleep(10 * time.Second)

	}
}
