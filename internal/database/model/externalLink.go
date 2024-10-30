package model

import (
	"time"

	"gorm.io/gorm"
)

type ExternalLink struct {
	ID         uint           `json:"-" gorm:"column:id;primaryKey"`
	ExternalId uint           `json:"id" gorm:"column:external_id"`
	SiteId     int            `json:"siteId" gorm:"column:site_id"`                                        // Site name from API
	Name       string         `json:"site" gorm:"column:name"`                                             // Site name from API
	Url        string         `json:"url" gorm:"column:url"`                                               // URL link from API
	Type       string         `json:"type" gorm:"column:type"`                                             // Type of link (e.g., "STREAMING")
	Language   string         `json:"language" gorm:"column:language"`                                     // Language from the api
	MediaId    uint           `json:"-" gorm:"column:media_id"`                                            // Foreign key to the media table
	CreatedAt  time.Time      `json:"createdAt" gorm:"column:created_at;type:timestamptz;"`                // Timestamps for DB
	UpdatedAt  time.Time      `json:"updatedAt" gorm:"column:updated_at;type:timestamptz;"`                // Timestamps for DB
	DeletedAt  gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deleted_at;type:timestamptz;index"` // Soft delete with timezone
}

func (ExternalLink) TableName() string {
	return "external_link" // Make sure GORM uses the correct table name
}
