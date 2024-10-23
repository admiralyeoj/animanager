package repository

import (
	"sync"

	aniListModel "github.com/admiralyeoj/anime-announcements/internal/aniListApi/model"
	"github.com/machinebox/graphql"
)

// AniListRepository defines the interface for AniList repository actions
type AniListRepository interface {
	GetUpcomingAnime(startDate, endDate string) (aniListModel.UpcomingAnime, error)
}

// aniListRepository is a concrete implementation of AniListRepository
type aniListRepository struct {
	graphqlClient *graphql.Client
}

// Use a pointer for the static instance
var instance *aniListRepository
var once sync.Once

// NewAniListRepository creates and returns a new instance of aniListRepository
func NewAniListRepository() AniListRepository {
	once.Do(func() {
		client := graphql.NewClient("https://graphql.anilist.co")
		instance = &aniListRepository{
			graphqlClient: client,
		}
	})

	return instance
}
