package cronJob

import (
	"fmt"
	"time"

	"github.com/admiralyeoj/animanager/internal/repository"
	"github.com/admiralyeoj/animanager/internal/service"
)

// TestCronJob struct implements CronJobInterface
type TestCronJob struct {
	FunctionName string
	Expression   string
	LastRun      *time.Time
	NextRun      *time.Time
}

// NewTestCronJob creates a new TestCronJob instance
func NewTestCronJob() *TestCronJob {
	return &TestCronJob{
		FunctionName: "testCommand",
		Expression:   "*/2 * * * *", // every 2 minutes
	}
}

func (j *TestCronJob) SetCronExpression(expression string) {
	j.Expression = expression
}

// CronExpression returns the cron expression for the job
func (j *TestCronJob) GetCronExpression() string {
	return j.Expression // Return the field CronExpression
}

// Handler executes the job logic
func (j *TestCronJob) Handler(srvs *service.Services, repos *repository.Repositories, params map[string]interface{}) error {
	fmt.Println("Running Test Job")

	return nil
}
