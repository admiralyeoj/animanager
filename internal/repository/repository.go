package repository

import (
	aniListRepo "github.com/admiralyeoj/animanager/internal/aniList/repository"
	blueSkyRepo "github.com/admiralyeoj/animanager/internal/blueSky/repository"
	"github.com/admiralyeoj/animanager/internal/config"
	dbRepos "github.com/admiralyeoj/animanager/internal/database/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	AniListRepo   aniListRepo.AniListRepository
	BlueSkyRepo   blueSkyRepo.BlueSkyRepository
	DatabaseRepos *dbRepos.DatabaseRepositories
	// Add other repositories here
}

func InitializeRepositories(cfg *config.Config, db *gorm.DB) *Repositories {
	return &Repositories{
		AniListRepo:   aniListRepo.NewAniListRepositories(),
		BlueSkyRepo:   blueSkyRepo.NewBlueSkyRepositories(cfg),
		DatabaseRepos: dbRepos.NewDatabaseRepositories(db),
		// Add other repositories here
	}
}
