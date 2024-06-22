// Package main is the entry point for the application
package main

import (
	"log"
	"net/http"

	"dispatch-auction/internal/handler"
)

func main() {
	log.Println("hello!")

	handler.SetupRESTHandlers()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
