package service

import (
	"database/sql"

	"github.com/admiralyeoj/anime-announcements/internal/aniListApi/service"
	"github.com/admiralyeoj/anime-announcements/internal/repository"
)

type Services struct {
	AniListSrv service.AniListService
	// Add other repositories here
}

func InitializeServices(repos *repository.Repositories, db *sql.DB) *Services {
	return &Services{
		AniListSrv: service.NewAniListService(*repos.DatabaseRepos),
		// Add other repositories here
	}
}
