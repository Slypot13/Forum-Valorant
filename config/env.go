package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Pas de fichier .env")
	}
}

func GetEnvWithDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}

	return value
}

func GetRequiredEnv(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists || value == "" {
		log.Fatalf("Variable manquante : %s", key)
	}

	return value
}