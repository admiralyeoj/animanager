package repository

import (
	"github.com/admiralyeoj/animanager/internal/database/model"
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
	// Step 1: Upsert for the parent (Media)
	err := mediaRepo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "external_id"}}, // Unique key to handle conflicts
		DoUpdates: clause.AssignmentColumns([]string{"site_url", "type", "format", "duration", "episodes", "cover_img", "banner_img", "updated_at"}),
	}).Omit("Title", "ExternalLinks").Create(&media).Error

	if err != nil {
		return err // Handle error appropriately
	}

	// Step 2: Check if MediaTitle is set in Media
	if media.Title.ID != 0 { // Ensure that MediaTitle is set (ID should not be zero)
		// update title
	}

	// if len(media.ExternalLinks) > 0 { // Ensure that MediaTitle is set (ID should not be zero)
	// 	// update links
	// }

	return nil
}
