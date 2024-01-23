package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
    // Load environment variables from .env file
    if err := godotenv.Load("src/cmd/main/.env"); err != nil {
        log.Println("Error loading .env file")
    }
}


func GetEnvVariable(key string) string {
	if key == "" {
		log.Fatalf("Key doens't exist")
	}
  
	return os.Getenv(key)
}