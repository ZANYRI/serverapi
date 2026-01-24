package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(name_key string) (result string, err error) {
	
	err = godotenv.Load("./env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	result = os.Getenv(name_key)

	return result, err
}
