package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load("app.env"); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}