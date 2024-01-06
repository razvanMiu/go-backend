package config

import (
	"log"

	"github.com/joho/godotenv"
)

func loadEnv() {
	// Load .env file
	envError := godotenv.Load()

	// If .env file is not found, ignore it
	if envError != nil {
		log.Fatal("Error loading .env file")
	}
}
