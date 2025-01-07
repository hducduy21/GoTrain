package main

import (
	"database/sql"
	"log"

	"github.com/hducduy21/gofull/cmd/api"
	"github.com/hducduy21/gofull/configs"
	"github.com/hducduy21/gofull/db"
)

func main() {
	// Code
	log.Println("Starting API server...")

	db, err := db.NewPQSQLStorage(configs.Envs.DbConnStr)
	if err != nil {
		log.Fatal(err)
	}

	pingStorage(&db)

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func pingStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
