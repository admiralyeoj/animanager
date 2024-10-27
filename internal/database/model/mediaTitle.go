package model

import "time"

type MediaTitle struct {
	ID        uint      `json:"-" gorm:"column:id;primaryKey"`
	English   string    `json:"english" gorm:"column:english"`
	MediaID   uint      `json:"-" gorm:"column:media_id;unique"` // Foreign key to Media
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (MediaTitle) TableName() string {
	return "media_title" // Make sure GORM uses the correct table name
}
