package service

import (
	"sync"

	bSkyRepo "github.com/admiralyeoj/animanager/internal/blueSky/repository"
	dbRepo "github.com/admiralyeoj/animanager/internal/database/repository"
)

// BlueSkyService defines the interface for Bluesky service actions
type BlueSkyService interface {
	// functions go here
}

// blueSkyService is a concrete implementation of BlueSkyService
type blueSkyService struct {
	bSkyRepository bSkyRepo.BlueSkyRepository
	dbRepositories dbRepo.DatabaseRepositories
}

// Use a pointer for the static instance
var instance *blueSkyService
var once sync.Once

// NewBlueSkyService returns the singleton instance of blueSkyService
func NewBlueSkyService(dbRepo dbRepo.DatabaseRepositories, bskyRepo bSkyRepo.BlueSkyRepository) BlueSkyService {
	return &blueSkyService{
		dbRepositories: dbRepo,
		bSkyRepository: bskyRepo,
	}
}
