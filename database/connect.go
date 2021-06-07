package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {

	db, err := sql.Open("postgres", "postgres://postgres:2400@localhost/ecomm?sslmode=disable")

	if err != nil {
		log.Panic("Database connection was not successful: ", err)
	}

	DB = db
}
