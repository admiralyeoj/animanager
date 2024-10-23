package config

import (
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// DatabaseConfig represents the configuration for the database
type DBConfig struct {
	Name     string `env:"DB_NAME" json:",omitempty"`
	Host     string `env:"DB_HOST, default=localhost" json:",omitempty"`
	User     string `env:"DB_USER" json:",omitempty"`
	Password string `env:"DB_PASSWORD" json:"-"` // ignored by zap's JSON formatter
	Port     string `env:"DB_PORT, default=5432" json:",omitempty"`
	URL      string `env:"DB_URL" json:",omitempty"`
}

func LoadConfig() (*DBConfig, error) {
	// Load the .env file, if it exists
	err := godotenv.Load() // It loads .env from the root directory by default
	if err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	// Parse the environment variables into the struct
	config := DBConfig{}
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// getEnv retrieves the value of an environment variable or returns the fallback value
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
