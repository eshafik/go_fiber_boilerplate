package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Define all environments here

var (
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
)

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}
	POSTGRES_HOST = getEnv("POSTGRES_HOST", "localhost")
	POSTGRES_PORT = getEnv("POSTGRES_PORT", "5432")
	POSTGRES_USER = getEnv("POSTGRES_USER", "postgres")
	POSTGRES_PASSWORD = getEnv("POSTGRES_PASSWORD", "postgres")
	POSTGRES_DB = getEnv("POSTGRES_DB", "postgres")
}
