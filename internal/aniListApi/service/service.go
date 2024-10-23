package service

import (
	"sync"

	"github.com/admiralyeoj/anime-announcements/internal/aniListApi/repository"
)

// AniListService defines the interface for AniList service actions
type AniListService interface {
	ImportUpcomingAnime(startDate, endDate string) error
}

// aniListService is a concrete implementation of AniListService
type aniListService struct {
	aniListRepository repository.AniListRepository // Embed the service
}

// Use a pointer for the static instance
var instance *aniListService
var once sync.Once

// NewAniListService returns the singleton instance of aniListService
func NewAniListService() AniListService {
	repo := repository.NewAniListRepository()

	return &aniListService{
		aniListRepository: repo,
	}
}