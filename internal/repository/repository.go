package repository

import (
	"database/sql"

	aniListRepo "github.com/admiralyeoj/anime-announcements/internal/aniListApi/repository"
	dbRepos "github.com/admiralyeoj/anime-announcements/internal/database/repository"
)

type Repositories struct {
	AniListRepo   aniListRepo.AniListRepository
	DatabaseRepos *dbRepos.DatabaseRepositories
	// Add other repositories here
}

func InitializeRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		AniListRepo:   aniListRepo.NewAniListRepositories(),
		DatabaseRepos: dbRepos.NewDatabaseRepositories(db),
		// Add other repositories here
	}
}
