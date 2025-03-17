package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetString(key string, fallback string) string {
	load()
	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return value
}

func load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
