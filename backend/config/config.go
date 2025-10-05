package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

}

func GetEnv(key string) string {
	return os.Getenv(key)
}