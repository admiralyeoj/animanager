package database

import (
	"log"
	"sync"

	"github.com/admiralyeoj/anime-announcements/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

// Connect initializes the database connection if not already created, otherwise returns the existing one.
func Connect(cfg *config.Config) (*gorm.DB, error) {
	var err error

	// Use sync.Once to ensure the connection is only initialized once
	once.Do(func() {
		dsn := cfg.DB.DSN // Your DSN should be defined in your config
		dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
			return
		}
	})

	// If an error occurred during initialization, return the error
	if err != nil {
		return nil, err
	}

	// Return the singleton instance
	return dbInstance, nil
}

// Close closes the database connection
func Close(db *gorm.DB) {
	sqlDB, err := db.DB() // Retrieve the underlying *sql.DB instance
	if err != nil {
		log.Fatalf("failed to get DB from GORM: %v", err)
	}

	// Close the database connection
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("failed to close the database connection: %v", err)
	}
}
