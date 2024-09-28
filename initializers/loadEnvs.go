package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	env := os.Getenv("GO_ENV")

	var err error
	// for production
	if env == "production" {
		err = godotenv.Load(".env.production")
		if err != nil {
			log.Fatalf("Error loading .env.production file: %v", err)
		}
		log.Println("Loaded environment variables from .env.production file")

	// for development
	} else {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
		log.Println("Loaded environment variables from .env file")
	}
}
