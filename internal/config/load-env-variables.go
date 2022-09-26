package config

import (
	"log"

	"github.com/joho/godotenv"
)

// ========== Load .env variables ==========
func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}
