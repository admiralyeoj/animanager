package repository

import (
	"database/sql"

	"github.com/admiralyeoj/anime-announcements/internal/aniListApi/model"
)

type MediaRepository interface {
	// Define your methods here
}

// mediaRepository is a concrete implementation of MediaRepository
type mediaRepository struct {
	db *sql.DB
}

func NewMediaRepository(db *sql.DB) MediaRepository {
	return &mediaRepository{
		db: db,
	}
}

func (mr *mediaRepository) Create(media *model.Media) error {
	return nil
}
