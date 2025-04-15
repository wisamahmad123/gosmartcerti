package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}