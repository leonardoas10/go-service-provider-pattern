package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Println("Error loading .env file")
    }
}


func GetEnvVariable(key string) string {
	if key == "" {
		log.Fatalf("Key doens't exist")
	}
  
	return os.Getenv(key)
}