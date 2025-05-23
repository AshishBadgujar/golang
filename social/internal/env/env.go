package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Load loads environment variables from .env file
func Load() error {
	return godotenv.Load()
}

// LoadFrom loads environment variables from a specific .env file
func LoadFrom(filename string) error {
	return godotenv.Load(filename)
}

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	num, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return num
}
