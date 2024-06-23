// Package main is the entry point for the application
package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq" // postgres driver

	"dispatch-auction/internal/handler"
	"dispatch-auction/internal/logic"
	"dispatch-auction/internal/share/util"
	"dispatch-auction/internal/storage"
)

const migrationFilePath = "migration_001.sql"

func main() {
	util.Init()

	connectionInfo := "user=postgres password=postgres host=127.0.0.1 port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		log.Fatal(err)
	}

	database := storage.New(db)
	err = database.Migrate(migrationFilePath)
	if err != nil {
		log.Fatal(err)
	}

	logic := logic.New(database)

	handler.SetupRESTHandlers(logic)

	log.Println("listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
