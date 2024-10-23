package service

import (
	"sync"

	"github.com/admiralyeoj/anime-announcements/internal/database/repository"
)

// DatabaseService defines the interface for Database service actions
type DatabaseService interface {
}

// databaseService is a concrete implementation of DatabaseService
type databaseService struct {
	databaseService repository.DatabaseRepository // Embed the service
}

// Use a pointer for the static instance
var instance *databaseService
var once sync.Once

func NewAniListService() DatabaseService {
	srv := repository.NewDatabaseRepository()

	return &databaseService{
		databaseService: srv,
	}
}
