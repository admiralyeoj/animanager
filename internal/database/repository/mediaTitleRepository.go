package repository

import (
	"github.com/admiralyeoj/anime-announcements/internal/aniListApi/model"
	"gorm.io/gorm"
)

type MediaTitleRepository interface {
	// Define your methods here
	Create(mediaId uint, title *model.MediaTitle) error
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
