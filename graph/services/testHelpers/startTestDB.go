package testhelpers

import (
	"database/sql"
	"log"
	"os"

	db "github/Martin-Martinez4/metube_backend/config"

	"github.com/joho/godotenv"
)

func StartTestDB() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TEST_DB_URL := os.Getenv("TEST_DB_URL")

	return db.GetDB("postgres", TEST_DB_URL)
}
