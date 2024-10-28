package model

import (
	"time"

	"gorm.io/gorm"
)

type Scheduler struct {
	ID             uint           `gorm:"primaryKey"`        // Primary key
	JobName        string         `gorm:"size:255;not null"` // Name of the job
	CronExpression string         `gorm:"size:255;not null"` // Cron expression for scheduling
	FunctionName   string         `gorm:"size:255;not null"` // Identifier of the function to execute
	IsActive       bool           `gorm:"default:true"`      // Enable or disable the job
	LastRun        *time.Time     `gorm:"type:timestamp"`    // Last execution time
	NextRun        *time.Time     `gorm:"type:timestamp"`    // Optional: next scheduled run time
	Params         []byte         `gorm:"type:jsonb"`        // Additional parameters in JSON format
	CreatedAt      time.Time      `gorm:"autoCreateTime"`    // Timestamp when job was created
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`    // Timestamp when job was last updated
	DeletedAt      gorm.DeletedAt `gorm:"index"`             // Soft delete, if needed
}

func (Scheduler) TableName() string {
	return "scheduler" // Make sure GORM uses the correct table name
}
