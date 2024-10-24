package database

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/admiralyeoj/anime-announcements/config"
	_ "github.com/lib/pq"
)

var (
	dbInstance *sql.DB
	once       sync.Once
)

// Connect initializes the database connection if not already created, otherwise returns the existing one.
func Connect(cfg *config.Config) (*sql.DB, error) {
	var err error

	// Use sync.Once to ensure the connection is only initialized once
	once.Do(func() {
		dbInstance, err = sql.Open("postgres", cfg.DB.DSN)
		if err != nil {
			return
		}

		dbInstance.SetMaxOpenConns(cfg.DB.MaxOpenConns)
		dbInstance.SetMaxIdleConns(cfg.DB.MaxIdleConns)

		duration, err := time.ParseDuration(cfg.DB.MaxIdleTime)
		if err != nil {
			return
		}

		dbInstance.SetConnMaxIdleTime(duration)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = dbInstance.PingContext(ctx)
	})

	// If an error occurred during initialization, return the error
	if err != nil {
		return nil, err
	}

	// Return the singleton instance
	return dbInstance, nil
}
