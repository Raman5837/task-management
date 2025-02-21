package utils

import (
	"strconv"
	"strings"

	"github.com/Raman5837/task-management/base/constants"
	_ "github.com/joho/godotenv/autoload" // Auto Load .env File
)

func IsProdEnv() bool {

	environment := strings.ToUpper(constants.GetEnv("ENVIRONMENT_NAME"))
	return strings.HasPrefix(environment, "PROD")

}

// Converts a string to boolean
func StringToBool(value string) bool {
	boolean, err := strconv.ParseBool(value)

	if err != nil {
		return false
	}
	return boolean
}

// Converts a string to an integer
func StringToInt(value string) *int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return nil
	}
	return &intValue
}

// Get actual value of the given pointer
func DereferencePointer(pointer *int, defaultValue int) int {

	if pointer != nil {
		return *pointer
	}

	return defaultValue
}
