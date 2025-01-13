package main

import (
	godotevn "github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotevn.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	username := os.Getenv("TADO_USERNAME")
	password := os.Getenv("TADO_PASSWORD")
	clientSecret := os.Getenv("TADO_CLIENT_SECRET")

	log.Printf("Username: %s\n", username)
	log.Printf("Password: %s\n", password)
	log.Printf("Client Secret: %s\n", clientSecret)
}
