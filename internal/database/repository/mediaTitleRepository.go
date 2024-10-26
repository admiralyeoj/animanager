package repository

import (
	"github.com/admiralyeoj/animanager/internal/database/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MediaTitleRepository interface {
	// Define your methods here
	Create(mediaId uint, title *model.MediaTitle) error
	UpdateOrCreate(mediaId uint, title *model.MediaTitle) error
}

// mediaRepository is a concrete implementation of MediaRepository
type mediaTitleRepository struct {
	db *gorm.DB
}

func NewMediaTitleRepository(db *gorm.DB) MediaTitleRepository {
	return &mediaTitleRepository{
		db: db,
	}
}

func (titleRepo *mediaTitleRepository) Create(mediaId uint, title *model.MediaTitle) error {
	title.MediaId = mediaId

	if err := titleRepo.db.Create(&title).Error; err != nil {
		return err // Return any error encountered during insertion
	}

	return nil
}

func (titleRepo *mediaTitleRepository) UpdateOrCreate(mediaId uint, title *model.MediaTitle) error {
	title.MediaId = mediaId

	tx := titleRepo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "media_id"}},           // Unique key to handle conflicts
		DoUpdates: clause.AssignmentColumns([]string{"english"}), // Fields to update
	}).Create(&title)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
