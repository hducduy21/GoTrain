package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPQSQLStorage(connStr string) (sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return *db, nil
}
