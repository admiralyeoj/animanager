package service

import (
	"sync"

	aniRepo "github.com/admiralyeoj/anime-announcements/internal/aniListApi/repository"
	dbRepo "github.com/admiralyeoj/anime-announcements/internal/database/repository"
)

// AniListService defines the interface for AniList service actions
type AniListService interface {
	ImportUpcomingAnime(startDate, endDate string) error
}

// aniListService is a concrete implementation of AniListService
type aniListService struct {
	aniListRepository aniRepo.AniListRepository // Embed the service
	dbRepositories    dbRepo.DatabaseRepositories
}

// Use a pointer for the static instance
var instance *aniListService
var once sync.Once

// NewAniListService returns the singleton instance of aniListService
func NewAniListService(dbRepo dbRepo.DatabaseRepositories) AniListService {
	return &aniListService{
		aniListRepository: aniRepo.NewAniListRepositories(),
		dbRepositories:    dbRepo,
	}
}
