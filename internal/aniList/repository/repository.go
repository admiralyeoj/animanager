package repository

import (
	aniListModel "github.com/admiralyeoj/anime-announcements/internal/database/model"
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

// NewAniListRepository creates and returns a new instance of aniListRepository
func NewAniListRepositories() AniListRepository {
	client := graphql.NewClient("https://graphql.anilist.co")
	return &aniListRepository{
		graphqlClient: client,
	}
}
