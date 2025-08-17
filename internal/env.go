package internal

import (
	"log"

	"github.com/joho/godotenv"
)

func GetProcessEnv() map[string]string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	return map[string]string{
		"PORT":     GetEnvReturnDefault("PORT", "3000"),
		"APP_NAME": GetEnvReturnDefault("APP_NAME", "default app"),
		"DB_HOST":  GetEnvReturnDefault("DB_HOST", "localhost"),
		"DB_PORT":  GetEnvReturnDefault("DB_PORT", "5432"),
		"DB_USER":  GetEnvReturnDefault("DB_USER", "postgres"),
		"DB_PASS":  GetEnvReturnDefault("DB_PASS", "postgres"),
		"DB_NAME":  GetEnvReturnDefault("DB_NAME", "postgres"),
		"ENV":      GetEnvReturnDefault("ENV", "dev"),
	}
}
