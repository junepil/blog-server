package main

import (
	"blog-api/internal/api"
	"blog-api/internal/config"
	"blog-api/internal/database"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	defer db.Close()

	log.Println("database connected successfully")

	router := api.NewRouter(db)
	log.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
