package service

import (
	aniList "github.com/admiralyeoj/animanager/internal/aniList/service"
	bluesky "github.com/admiralyeoj/animanager/internal/blueSky/service"
	"github.com/admiralyeoj/animanager/internal/repository"
	"gorm.io/gorm"
)

type Services struct {
	AniListSrv aniList.AniListService
	BlueSkySrv bluesky.BlueSkyService
	// Add other repositories here
}

func InitializeServices(repos *repository.Repositories, db *gorm.DB) *Services {
	return &Services{
		AniListSrv: aniList.NewAniListService(repos.DatabaseRepos),
		BlueSkySrv: bluesky.NewBlueSkyService(repos.DatabaseRepos, repos.BlueSkyRepo),
		// Add other repositories here
	}
}
