package constants

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload" // Auto Load .env File
)

// Utility function to get `env` value
func GetEnv(key string) string {

	if _, notFound := os.Stat(".env"); notFound != nil {
		fmt.Printf(".env file not found --> %v", notFound.Error())
	}

	// Load .env File
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error while loading .env file", err)
	}
	// Return the value of the variable
	return os.Getenv(key)
}
