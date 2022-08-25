package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DB() *sql.DB {
	connection := os.Getenv("DB_CONNECTION")
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	return db
}
