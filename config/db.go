package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=p@ssw0rd dbname=cmis sslmode=disable")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database not responding:", err)
	}

	log.Println("Connected to DB")
}
