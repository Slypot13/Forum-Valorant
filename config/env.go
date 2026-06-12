package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

// LoadEnv charge le fichier .env.
func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Pas de fichier .env")
	}
}

// GetEnvWithDefault récupère une variable ou sa valeur par défaut.
func GetEnvWithDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}

	return value
}

// GetRequiredEnv récupère une variable obligatoire.
func GetRequiredEnv(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists || value == "" {
		log.Fatalf("Variable manquante : %s", key)
	}

	return value
}