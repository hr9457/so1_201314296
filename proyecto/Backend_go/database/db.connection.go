package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetDB() (*sql.DB, error) {

	// variables de entorno
	var err = godotenv.Load("./.env")

	// verificcion de error con la conexion a db
	if err != nil {
		log.Fatal("Error loading .env")
	}

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("USER_DB"),
		os.Getenv("PASS_DB"),
		os.Getenv("HOTS_DB"),
		os.Getenv("PORT_DB"),
		os.Getenv("NAME_DB"))

	return sql.Open("mysql", connectionString)
}

func TestConnection() {
	db, err := GetDB()
	if err != nil {
		log.Printf("Error with database DBSO1" + err.Error())
		return
	} else {
		err = db.Ping()
		if err != nil {
			log.Printf("Error making connection to DB. Please check credentials. The error is: " + err.Error())
			return
		}
		log.Printf("Connect with database DBSO1")
	}
}
