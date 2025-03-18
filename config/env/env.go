package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetString(key string, fallback string) string {

	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return value
}

func GetInt(key string, fallback int) int {

	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return intValue
}

func GetBool(key string, fallback bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}

	return boolValue
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
