package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	testEnv := os.Getenv("PORT")
	if testEnv == "" {
		log.Fatal("PORT environment variable is not set")
	}
}
