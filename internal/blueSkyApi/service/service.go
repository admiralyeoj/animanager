package service

import (
	"sync"

	dbRepo "github.com/admiralyeoj/anime-announcements/internal/database/repository"
)

// AniListService defines the interface for AniList service actions
type AniListService interface {
	// functions go here
}

// aniListService is a concrete implementation of AniListService
type aniListService struct {
	dbRepositories dbRepo.DatabaseRepositories
}

// Use a pointer for the static instance
var instance *aniListService
var once sync.Once

// NewAniListService returns the singleton instance of aniListService
func NewAniListService(dbRepo dbRepo.DatabaseRepositories) AniListService {
	return &aniListService{
		dbRepositories: dbRepo,
	}
}
