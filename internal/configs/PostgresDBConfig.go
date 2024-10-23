package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DatabaseConfig represents the configuration for the database
type PostgresDBConfig struct {
	URL      string
	Username string
	Password string
	Port     string
	Name     string
}

// LoadConfig reads the configuration from the .env file
func LoadConfig() (*PostgresDBConfig, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables.")
	}

	// Create the database configuration struct
	config := &PostgresDBConfig{
		URL:      getEnv("DATABASE_URL", "default_database_url"),
		Username: getEnv("DATABASE_USERNAME", ""),
		Password: getEnv("DATABASE_PASSWORD", ""),
		Port:     getEnv("DATABASE_PORT", "5432"),
		Name:     getEnv("DATABASE_NAME", ""),
	}

	return config, nil
}

// getEnv retrieves the value of an environment variable or returns the fallback value
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
