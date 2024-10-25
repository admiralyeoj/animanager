package repository

import (
	"github.com/admiralyeoj/anime-announcements/internal/aniListApi/model"
	"gorm.io/gorm"
)

type AiringScheduleRepository interface {
	// Define your methods here
	Create(mediaId uint, schedules *model.AiringSchedule) error
}

// mediaRepository is a concrete implementation of MediaRepository
type airingScheduleRepository struct {
	db *gorm.DB
}

func NewAiringScheduleRepository(db *gorm.DB) AiringScheduleRepository {
	return &airingScheduleRepository{
		db: db,
	}
}

func (linksRepo *airingScheduleRepository) Create(mediaId uint, schedule *model.AiringSchedule) error {
	schedule.MediaId = mediaId

	if err := linksRepo.db.Create(&schedule).Error; err != nil {
		return err // Return any error encountered during insertion
	}

	return nil
}
