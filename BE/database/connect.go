package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/ocg-database")
	if err != nil {
		log.Printf("Error %s when opening DBn", err)
		return nil, err
	}
	DB = db

	return db, nil
}