package model

import "time"

type SocialPost struct {
	ID               uint      `json:"-" gorm:"column:id;primaryKey"`
	PostId           string    `json:"post_id" gorm:"column:post_id"` // English title from API
	AiringScheduleId uint      `json:"-" gorm:"column:airing_schedule_id;not null"`
	CreatedAt        time.Time `json:"createdAt" gorm:"column:created_at"` // Timestamps for DB
	UpdatedAt        time.Time `json:"updatedAt" gorm:"column:updated_at"` // Timestamps for DB
}

func (SocialPost) TableName() string {
	return "social_post" // Make sure GORM uses the correct table name
}
