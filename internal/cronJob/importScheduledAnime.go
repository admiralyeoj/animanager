package cronJob

import (
	"fmt"
	"time"

	AniListSrv "github.com/admiralyeoj/animanager/internal/aniList/service"
)

// ImportScheduledAnimeCronJob struct implements CronJobInterface
type ImportScheduledAnimeCronJob struct {
	FunctionName string
	Expression   string
	LastRun      *time.Time
	NextRun      *time.Time
	aniListSrv   *AniListSrv.AniListService
}

// NewImportScheduledAnimeCronJob creates a new ImportScheduledAnimeCronJob instance
func NewImportScheduledAnimeCronJob(aniListSrv *AniListSrv.AniListService) *ImportScheduledAnimeCronJob {
	return &ImportScheduledAnimeCronJob{
		FunctionName: "importScheduledAnime",
		aniListSrv:   aniListSrv,
	}
}

func (j *ImportScheduledAnimeCronJob) SetCronExpression(expression string) {
	j.Expression = expression
}

// CronExpression returns the cron expression for the job
func (j *ImportScheduledAnimeCronJob) GetCronExpression() string {
	return j.Expression // Return the field CronExpression
}

// Handler executes the job logic
func (j *ImportScheduledAnimeCronJob) Handler(params map[string]interface{}, args ...interface{}) error {

	format := "01/02/2006"

	startDate := time.Now().Format(format)
	date, _ := time.Parse(format, startDate)
	endDate := date.AddDate(0, 0, 1).Format(format)

	err := (*j.aniListSrv).ImportUpcomingAnime(startDate, endDate)
	if err != nil {
		return fmt.Errorf("error importing anime: %w", err)
	}

	fmt.Println("Anime successfully imported")
	return nil
}
