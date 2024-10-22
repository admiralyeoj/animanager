package aniListApi

import (
	"time"

	"github.com/machinebox/graphql"
	// "github.com/admiralyeoj/pokedexcli/internal/pokeCache"
)

// Client -
type Client struct {
	// cache      pokeCache.Cache
	graphqlClient *graphql.Client
}

// NewClient -
func NewClient(cacheInterval time.Duration) Client {
	client := graphql.NewClient("https://graphql.anilist.co")

	return Client{
		// cache: pokeCache.NewCache(cacheInterval),
		graphqlClient: client,
	}
}
