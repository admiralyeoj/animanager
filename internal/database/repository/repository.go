package repository

import (
	"sync"
)

// DatabaseRepository defines the interface for database repository actions
type DatabaseRepository interface {
}

// databaseRepository is a concrete implementation of DatabaseRepository
type databaseRepository struct {
}

// Use a pointer for the static instance
var instance *databaseRepository
var once sync.Once

// NewAniListRepository creates and returns a new instance of aniListRepository
func NewDatabaseRepository() DatabaseRepository {
	once.Do(func() {
		instance = &databaseRepository{}
	})

	return instance
}
