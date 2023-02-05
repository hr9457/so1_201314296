package database

import (
	"fmt"
	"log"

	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetDB() string {

	var err = godotenv.Load("./.env")

	// verificcion de error con la conexion a db
	if err != nil {
		log.Fatal("Error loading .env")
	}

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
	return connectionString

	// return sql.Open("mysql", connectionString)
}
