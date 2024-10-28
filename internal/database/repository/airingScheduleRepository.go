package repository

import (
	"errors"
	"time"

	"github.com/admiralyeoj/animanager/internal/database/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AiringScheduleRepository interface {
	// Define your methods here
	Create(mediaId uint, schedules *model.AiringSchedule) error
	UpdateOrCreate(mediaId uint, schedule *model.AiringSchedule) error
	GetNextNotAnnounced() (*model.AiringSchedule, error)
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

	err := airing.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "external_id"}},                                               // Unique key to handle conflicts
		DoUpdates: clause.AssignmentColumns([]string{"airing_at", "episode", "media_id", "updated_at"}), // Fields to update
	}).Omit("Media").Create(&schedule).Error

	if err != nil {
		return err
	}

	return nil
}

func (airing *airingScheduleRepository) GetNextNotAnnounced() (*model.AiringSchedule, error) {

	var result *model.AiringSchedule

	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()

	// Fetch AiringSchedule records without a corresponding SocialPost by airing_schedule_id
	err := airing.db.
		Not("id IN (?)", airing.db.Model(model.SocialPost{}).Select("airing_schedule_id")).
		Where("airing_at > ?", midnight).
		Where("airing_at < ?", now.Unix()).
		Preload("Media.Title").
		Preload("Media.ExternalLinks").
		Order("airing_at ASC").
		First(&result).Error

	// Only return an error if it's not ErrRecordNotFound
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return result, nil
}
