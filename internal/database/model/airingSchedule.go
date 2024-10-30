package model

import (
	"time"

	"gorm.io/gorm"
)

type AiringSchedule struct {
	ID          uint           `json:"-" gorm:"column:id;primaryKey"`                                       // Primary key for the airing schedule
	ExternalId  int            `json:"id" gorm:"column:external_id;unique"`                                 // Unique external ID corresponding to Anilist
	AiringAt    int64          `json:"airingAt" gorm:"column:airing_at"`                                    // Airing time
	Episode     int            `json:"episode" gorm:"column:episode"`                                       // Episode number
	MediaId     uint           `json:"mediaId" gorm:"column:media_id"`                                      // Foreign key to the media table
	Media       Media          `json:"media,omitempty" gorm:"foreignKey:MediaId"`                           // Related media data
	SocialPosts []SocialPost   `json:"socialPosts,omitempty" gorm:"foreignKey:AiringScheduleId"`            // Optional social posts
	CreatedAt   time.Time      `json:"createdAt" gorm:"column:created_at;type:timestamptz"`                 // Timestamp for creation
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"column:updated_at;type:timestamptz"`                 // Timestamp for update
	DeletedAt   gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deleted_at;type:timestamptz;index"` // Soft delete with timezone
}

func (AiringSchedule) TableName() string {
	return "airing_schedule" // Make sure GORM uses the correct table name
}
