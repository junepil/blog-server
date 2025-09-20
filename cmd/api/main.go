package main

import (
	"blog-api/internal/api"
	"blog-api/internal/database"
	"blog-api/internal/s3"
	"log"
	"net/http"
)

func main() {
	database.Connect()

	uploader, err := s3.NewUploader()
	if err != nil {
		log.Fatalf("could not create S3 uploader: %v", err)
	}

	router := api.NewRouter(database.DB, uploader)
	log.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
