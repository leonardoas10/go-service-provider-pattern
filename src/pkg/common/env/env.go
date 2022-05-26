package env

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func GetEnvVariable(key string) string {
	if key == "" {
		log.Fatalf("Key doens't exist")
	}
  
	return os.Getenv(key)
}