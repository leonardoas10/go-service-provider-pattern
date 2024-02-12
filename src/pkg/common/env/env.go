package env

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)


func init() {
    // Get the path to the project root directory
    rootDir := findRootDir()
	envFile := os.Getenv("APP_ENV")
	if envFile == "" {
		envFile = ".env" // default to regular environment file
	}

    // Construct the path to the .env file in the root directory
    pathToEnvFile := filepath.Join(rootDir, envFile)

    fmt.Println("pathToEnvFile ==> ", pathToEnvFile)

    // Load environment variables from the .env file
    if err := godotenv.Load(pathToEnvFile); err != nil {
        log.Printf("Error loading %s file: %v", envFile, err)
    }
}

func findRootDir() string {
    // Start from the current directory
    currentDir, err := os.Getwd()
    if err != nil {
        log.Fatalf("Error getting current directory: %v", err)
    }

    // Navigate up the directory tree until we find the root
    for {
        if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
            // Found the root directory
            return currentDir
        }

        // Go up one directory
        parentDir := filepath.Dir(currentDir)

        // If we're already at the root directory (no parent), break the loop
        if parentDir == currentDir {
            break
        }

        // Update current directory to the parent directory
        currentDir = parentDir
    }

    // If we couldn't find the root directory, return an empty string
    return ""
}


func GetEnvVariable(key string) string {
	if key == "" {
		log.Fatalf("Key doens't exist")
	}
  
	return os.Getenv(key)
}