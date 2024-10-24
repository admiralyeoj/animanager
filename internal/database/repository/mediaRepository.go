package repository

import (
	"database/sql"

	"github.com/admiralyeoj/anime-announcements/internal/aniListApi/model"
)

type MediaRepository interface {
	// Define your methods here
	Create(media *model.Media) error
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

func (med *mediaRepository) Create(media *model.Media) error {
	query := "INSERT INTO media (external_id, site_url, type_id, format_id, duration, episodes, cover_img, banner_img) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	err := med.db.QueryRow(query, media.ExternalId, media.SiteUrl, 1, 1, media.Duration, media.Episodes, media.CoverImage, media.BannerImage).Scan(&media.Id)

	if err != nil {
		return err
	}
	return nil
}
