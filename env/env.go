package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(provider string) (result string, err error) {

	err = godotenv.Load("./env/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	result = os.Getenv(provider)

	return result, err
}
