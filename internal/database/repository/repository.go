package repository

import (
	"sync"

	"gorm.io/gorm"
)

// DatabaseRepositories holds instances of different repositories
type DatabaseRepositories struct {
	Media          MediaRepository
	AiringSchedule AiringScheduleRepository
	MediaTitle     MediaTitleRepository
	ExternalLinks  ExternalLinksRepository
	SocialPost     SocialPostRepository
	Scheduler      SchedulerRepository
}

// Use a pointer for the static instance
var instance *DatabaseRepositories
var once sync.Once

// NewDatabaseRepositories creates and returns a new instance of DatabaseRepositories
func NewDatabaseRepositories(db *gorm.DB) *DatabaseRepositories {
	once.Do(func() {
		// Initialize the repositories

		instance = &DatabaseRepositories{
			Media:          NewMediaRepository(db),
			AiringSchedule: NewAiringScheduleRepository(db),
			MediaTitle:     NewMediaTitleRepository(db),
			ExternalLinks:  NewExternalLinksRepository(db),
			SocialPost:     NewSocialPostRepository(db),
			Scheduler:      NewSchedulerRepository(db),
		}
	})

	return instance
}
