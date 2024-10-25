package service

import (
	"github.com/admiralyeoj/anime-announcements/internal/aniListApi/service"
	"github.com/admiralyeoj/anime-announcements/internal/repository"
	"gorm.io/gorm"
)

type Services struct {
	AniListSrv service.AniListService
	// Add other repositories here
}

func InitializeServices(repos *repository.Repositories, db *gorm.DB) *Services {
	return &Services{
		AniListSrv: service.NewAniListService(*repos.DatabaseRepos),
		// Add other repositories here
	}
}
