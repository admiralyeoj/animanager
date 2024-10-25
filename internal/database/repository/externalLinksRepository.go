package repository

import (
	"errors"

	"github.com/admiralyeoj/anime-announcements/internal/aniListApi/model"
	"gorm.io/gorm"
)

type ExternalLinksRepository interface {
	// Define your methods here
	Create(mediaId uint, links *[]model.ExternalLink) error
}

// mediaRepository is a concrete implementation of MediaRepository
type externalLinksRepository struct {
	db *gorm.DB
}

func NewExternalLinksRepository(db *gorm.DB) ExternalLinksRepository {
	return &externalLinksRepository{
		db: db,
	}
}

func (linksRepo *externalLinksRepository) Create(mediaId uint, links *[]model.ExternalLink) error {
	for i := range *links {
		(*links)[i].MediaId = mediaId // Set the foreign key for each link
	}

	if len(*links) == 0 {
		return errors.New("no links to insert")
	}

	if err := linksRepo.db.Create(&links).Error; err != nil {
		return err // Return any error encountered during insertion
	}

	return nil
}
