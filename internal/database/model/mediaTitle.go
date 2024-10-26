package model

import "time"

type MediaTitle struct {
	ID        uint      `json:"-" gorm:"column:id;primaryKey"`
	English   string    `json:"english"` // English title from API
	MediaId   uint      `json:"-" gorm:"column:media_id;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"` // Timestamps for DB
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"` // Timestamps for DB
}

func (MediaTitle) TableName() string {
	return "media_title" // Make sure GORM uses the correct table name
}
