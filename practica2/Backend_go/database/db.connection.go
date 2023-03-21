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
	var err = godotenv.Load(".env")

	// verificcion de error con la conexion a db
	if err != nil {
		log.Fatal("Error loading .env")
	}

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("userDB"),
		os.Getenv("passDB"),
		os.Getenv("hostDB"),
		os.Getenv("portDB"),
		os.Getenv("nameDB"))

	return sql.Open("mysql", connectionString)
}

func TestConnection() {
	db, err := GetDB()
	if err != nil {
		log.Printf("Error with database DBSO1" + err.Error())
		return
	} else {
		err = db.Ping()
		fmt.Println(err)
		if err != nil {
			log.Printf("Error making connection to DB. The error is: " + err.Error())
			return
		}
		log.Printf("Connect with database DBSO1")
	}
}
