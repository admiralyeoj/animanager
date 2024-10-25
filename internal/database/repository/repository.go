package repository

import (
	"database/sql"
	"sync"
)

// DatabaseRepositories holds instances of different repositories
type DatabaseRepositories struct {
	Media MediaRepository
}

// Use a pointer for the static instance
var instance *DatabaseRepositories
var once sync.Once

// NewDatabaseRepositories creates and returns a new instance of DatabaseRepositories
func NewDatabaseRepositories(db *sql.DB) *DatabaseRepositories {
	once.Do(func() {
		// Initialize the repositories
		mediaRepo := NewMediaRepository(db)

		instance = &DatabaseRepositories{
			Media: mediaRepo, // This now matches the MediaRepository type
			// Initialize other repositories here as needed
		}
	})

	return instance
}
