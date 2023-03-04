package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDB() *sql.DB {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if db == nil {

		DB_URL := os.Getenv("DB_URL")

		db, err = sql.Open("postgres", DB_URL)
		if err != nil {
			panic(err)
		}

		if err = db.Ping(); err != nil {
			panic(err)
		}
		fmt.Println("You connected to your database.")
	}

	return db

}

func goDotEnvVariable(s string) {
	panic("unimplemented")
}
