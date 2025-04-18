package config

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

func ConnectDB() {
    var err error
    DB, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=yourpassword dbname=construction_app sslmode=disable")
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Database not responding:", err)
    }

    log.Println("Connected to DB")
}
