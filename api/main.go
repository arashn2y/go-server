package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/arashn2y/go-server/handlers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment")
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
