package cronJob

import (
	"fmt"
	"time"

	blueSkySrv "github.com/admiralyeoj/animanager/internal/blueSky/service"
)

// AnnounceNewAnimeCronJob struct implements CronJobInterface
type AnnounceNewAnimeCronJob struct {
	FunctionName string
	Expression   string
	LastRun      *time.Time
	NextRun      *time.Time
	blueSkySrv   *blueSkySrv.BlueSkyService
}

// NewAnnounceNewAnimeCronJob creates a new AnnounceNewAnimeCronJob instance
func NewAnnounceNewAnimeCronJob(blueSkySrv *blueSkySrv.BlueSkyService) *AnnounceNewAnimeCronJob {
	return &AnnounceNewAnimeCronJob{
		FunctionName: "testCommand",
		blueSkySrv:   blueSkySrv,
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
func (j *AnnounceNewAnimeCronJob) Handler(params map[string]interface{}, args ...interface{}) error {
	if err := (*j.blueSkySrv).AnnounceAiringAnime(); err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
