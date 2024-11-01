package repository

import (
	"github.com/admiralyeoj/animanager/internal/database/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialPostRepository interface {
	// Define your methods here
	Create(AiringScheduleId uint, post *model.SocialPost) error
	UpdateOrCreate(AiringScheduleId uint, post *model.SocialPost) error
}

// socialpostRepository is a concrete implementation of SocialPostRepository
type socialpostRepository struct {
	db *gorm.DB
}

func NewSocialPostRepository(db *gorm.DB) SocialPostRepository {
	return &socialpostRepository{
		db: db,
	}
}

func (postRepo *socialpostRepository) Create(AiringScheduleId uint, post *model.SocialPost) error {
	post.AiringScheduleId = AiringScheduleId

	if err := postRepo.db.Create(&post).Error; err != nil {
		return err // Return any error encountered during insertion
	}

	return nil
}

func (postRepo *socialpostRepository) UpdateOrCreate(AiringScheduleId uint, post *model.SocialPost) error {
	post.AiringScheduleId = AiringScheduleId

	tx := postRepo.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "airing_schedule_id"}},               // Unique key to handle conflicts
		DoUpdates: clause.AssignmentColumns([]string{"post_id", "updated_at"}), // Fields to update
	}).Create(&post)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
