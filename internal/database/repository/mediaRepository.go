package repository

import (
	"github.com/admiralyeoj/anime-announcements/internal/database/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MediaRepository interface {
	// Define your methods here
	Create(media *model.Media) error
	UpdateOrCreate(media *model.Media) error
}

// mediaRepository is a concrete implementation of MediaRepository
type mediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) MediaRepository {
	return &mediaRepository{
		db: db,
	}
}

func (mediaRepo *mediaRepository) Create(media *model.Media) error {
	if err := mediaRepo.db.Create(&media).Error; err != nil {
		return err // Return any error encountered during insertion
	}

	return nil
}

func (mediaRepo *mediaRepository) UpdateOrCreate(media *model.Media) error {
	tx := mediaRepo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "external_id"}},                                                                              // Unique key to handle conflicts
		DoUpdates: clause.AssignmentColumns([]string{"site_url", "type", "format", "duration", "episodes", "cover_img", "banner_img"}), // Fields to update
	}).Create(&media)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
