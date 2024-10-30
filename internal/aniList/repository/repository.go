package repository

import (
	aniListModel "github.com/admiralyeoj/animanager/internal/database/model"
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

func NewAniListRepositories() AniListRepository {
	client := graphql.NewClient("https://graphql.anilist.co")
	return &aniListRepository{ // Return a pointer to aniListRepository
		graphqlClient: client,
	}
}
