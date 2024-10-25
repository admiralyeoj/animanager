package repository

import (
	"github.com/admiralyeoj/anime-announcements/internal/aniListApi/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AiringScheduleRepository interface {
	// Define your methods here
	Create(mediaId uint, schedules *model.AiringSchedule) error
	UpdateOrCreate(mediaId uint, schedule *model.AiringSchedule) error
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

func (airing *airingScheduleRepository) Create(mediaId uint, schedule *model.AiringSchedule) error {
	schedule.MediaId = mediaId

	if err := airing.db.Create(&schedule).Error; err != nil {
		return err // Return any error encountered during insertion
	}

	return nil
}

func (airing *airingScheduleRepository) UpdateOrCreate(mediaId uint, schedule *model.AiringSchedule) error {
	schedule.MediaId = mediaId

	tx := airing.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "external_id"}},                                 // Unique key to handle conflicts
		DoUpdates: clause.AssignmentColumns([]string{"airing_at", "episode", "media_id"}), // Fields to update
	}).Create(&schedule)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
