package model

import (
	"time"

	"gorm.io/gorm"
)

type SocialPost struct {
	ID               uint           `json:"-" gorm:"column:id;primaryKey"`                                       // auto incrementing number
	PostId           string         `json:"post_id" gorm:"column:post_id"`                                       // social id for the post
	AiringScheduleId uint           `json:"-" gorm:"column:airing_schedule_id;not null"`                         // external aring id
	CreatedAt        time.Time      `json:"createdAt" gorm:"column:created_at;type:timestamptz;"`                // Timestamps for DB
	UpdatedAt        time.Time      `json:"updatedAt" gorm:"column:updated_at;type:timestamptz;"`                // Timestamps for DB
	DeletedAt        gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deleted_at;type:timestamptz;index"` // Soft delete with timezone

}

func (SocialPost) TableName() string {
	return "social_post" // Make sure GORM uses the correct table name
}
