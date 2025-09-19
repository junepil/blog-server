package main

import (
	"blog-api/internal/api"
	"blog-api/internal/database"
	"log"
	"net/http"
)

func main() {
	database.Connect()

	router := api.NewRouter(database.DB)
	log.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
