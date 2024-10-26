package service

import (
	"github.com/admiralyeoj/animanager/internal/aniList/service"
	"github.com/admiralyeoj/animanager/internal/repository"
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
