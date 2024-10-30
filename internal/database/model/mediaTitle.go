package model

import (
	"time"

	"gorm.io/gorm"
)

type MediaTitle struct {
	ID        uint           `json:"-" gorm:"column:id;primaryKey"`
	English   string         `json:"english" gorm:"column:english"`
	MediaID   uint           `json:"-" gorm:"column:media_id;unique"` // Foreign key to Media
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at;type:timestamptz;"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;type:timestamptz;"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deleted_at;type:timestamptz;index"` // Soft delete with timezone

}

func (MediaTitle) TableName() string {
	return "media_title" // Make sure GORM uses the correct table name
}
