package model

import (
	"time"

	"gorm.io/gorm"
)

type Media struct {
	ID            uint           `json:"-" gorm:"column:id;primaryKey"`
	ExternalId    uint           `json:"id" gorm:"column:external_id"`                                        // External ID for the media (from Anilist)
	SiteUrl       string         `json:"siteUrl" gorm:"column:site_url"`                                      // Media site URL
	Type          string         `json:"type" gorm:"column:type"`                                             // Media type (e.g., "ANIME") — map to type in DB
	Format        string         `json:"format" gorm:"column:format"`                                         // Format (e.g., "ONA") — map to format in DB
	Duration      int            `json:"duration" gorm:"column:duration"`                                     // Duration of the media
	Episodes      int            `json:"episodes" gorm:"column:episodes"`                                     // Number of episodes
	CoverImage    CoverImg       `json:"coverImage" gorm:"embedded"`                                          // Store the cover image URL in DB
	BannerImage   string         `json:"bannerImage" gorm:"column:banner_img"`                                // Store banner image URL in DB (if needed)
	Title         MediaTitle     `json:"title" gorm:"foreignKey:MediaID;references:ID"`                       // Foreign key setup for MediaTitle
	ExternalLinks []ExternalLink `json:"externalLinks,omitempty"`                                             // Ignore in DB (can store in another table if needed)
	CreatedAt     time.Time      `json:"createdAt" gorm:"column:created_at;type:timestamptz;"`                // Timestamps for DB
	UpdatedAt     time.Time      `json:"updatedAt" gorm:"column:updated_at;type:timestamptz;"`                // Timestamps for DB
	DeletedAt     gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deleted_at;type:timestamptz;index"` // Soft delete with timezone

}

type CoverImg struct {
	Large string `json:"large" gorm:"column:cover_img"` // Cover image URL from API
}
