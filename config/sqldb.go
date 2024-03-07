package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDB(driverName string, sqlURI string) *sql.DB {

	if db == nil {

		_db, err := sql.Open(driverName, sqlURI)
		db = _db
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
