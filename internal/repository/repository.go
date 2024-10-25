package repository

import (
	aniListRepo "github.com/admiralyeoj/anime-announcements/internal/aniListApi/repository"
	dbRepos "github.com/admiralyeoj/anime-announcements/internal/database/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	AniListRepo   aniListRepo.AniListRepository
	DatabaseRepos *dbRepos.DatabaseRepositories
	// Add other repositories here
}

func InitializeRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		AniListRepo:   aniListRepo.NewAniListRepositories(),
		DatabaseRepos: dbRepos.NewDatabaseRepositories(db),
		// Add other repositories here
	}
}
