package cronJob

import (
	"fmt"
	"time"
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
func (j *TestCronJob) Handler(params map[string]interface{}, args ...interface{}) error {
	fmt.Println("Running Test Job")

	return nil
}
