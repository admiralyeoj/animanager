package cronJob

import (
	"fmt"
	"time"

	dbRepos "github.com/admiralyeoj/animanager/internal/database/repository"
)

// AnnounceNewAnimeCronJob struct implements CronJobInterface
type AnnounceNewAnimeCronJob struct {
	FunctionName string
	Expression   string
	LastRun      *time.Time
	NextRun      *time.Time
	dbRepo       *dbRepos.DatabaseRepositories
}

// NewAnnounceNewAnimeCronJob creates a new AnnounceNewAnimeCronJob instance
func NewAnnounceNewAnimeCronJob(dbRepo *dbRepos.DatabaseRepositories) *AnnounceNewAnimeCronJob {
	return &AnnounceNewAnimeCronJob{
		FunctionName: "testCommand",
		dbRepo:       dbRepo,
	}
}

func (j *AnnounceNewAnimeCronJob) SetCronExpression(expression string) {
	j.Expression = expression
}

// CronExpression returns the cron expression for the job
func (j *AnnounceNewAnimeCronJob) GetCronExpression() string {
	return j.Expression // Return the field CronExpression
}

// Handler executes the job logic
func (j *AnnounceNewAnimeCronJob) Handler(params map[string]interface{}) error {
	fmt.Println("Running Test Job")

	return nil
}
