package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func loadEnvFile() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Success load .env file")
}
