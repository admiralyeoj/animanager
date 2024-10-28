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

/* https://pkg.go.dev/github.com/robfig/cron@v1.2.0#section-readme
+------------------+------------+--------------------------+-------------------------------+
| Field Name       | Mandatory? | Allowed Values           | Allowed Special Characters    |
+------------------+------------+--------------------------+-------------------------------+
| Seconds          | Yes        | 0-59                     | * / , -                       |
| Minutes          | Yes        | 0-59                     | * / , -                       |
| Hours            | Yes        | 0-23                     | * / , -                       |
| Day of Month     | Yes        | 1-31                     | * / , - ?                     |
| Month            | Yes        | 1-12 or JAN-DEC          | * / , -                       |
| Day of Week      | Yes        | 0-6 or SUN-SAT           | * / , - ?                     |
+------------------+------------+--------------------------+-------------------------------+
*/
