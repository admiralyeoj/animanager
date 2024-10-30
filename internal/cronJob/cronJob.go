package cronJob

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/admiralyeoj/animanager/internal/database/model"
	"github.com/admiralyeoj/animanager/internal/repository"
	"github.com/robfig/cron"
)

type CronJobInterface interface {
	GetCronExpression() string
	SetCronExpression(string)
	Handler(repos *repository.Repositories, params map[string]interface{}) error
}

// InitializeCronJobs loads jobs from the database and sets up the cron scheduler
func InitializeCronJobs(repos *repository.Repositories) *cron.Cron {
	// Initialize cron scheduler
	c := cron.New()

	// Retrieve all active cron jobs from the database
	activeJobs, err := repos.DatabaseRepos.Scheduler.GetActiveJobs()
	if err != nil {
		log.Fatalf("Failed to load jobs: %v", err)
	}

	// Create a mapping of job types to constructors
	jobMap := map[string]func() CronJobInterface{
		"test": func() CronJobInterface { return NewTestCronJob() },
	}

	// Convert database jobs to CronJob structs and register with cron
	for _, activeJob := range activeJobs {

		fmt.Println(activeJob.JobName)
		jobConstructor, exists := jobMap[activeJob.FunctionName]
		if !exists {
			log.Printf("No function mapped for job %s, skipping...\n", activeJob.JobName)
			continue
		}

		cronJob := jobConstructor()
		cronJob.SetCronExpression(activeJob.CronExpression) // Set the cron expression from the database

		var params map[string]interface{}
		// Unmarshal params from JSON
		if activeJob.Params != nil {
			if err := json.Unmarshal(activeJob.Params, &params); err != nil {
				log.Printf("Error unmarshalling params for job %s: %v\n", activeJob.JobName, err)
				return nil // Early exit on error
			}
		}

		// Schedule the job based on the cron expression from the database
		err := c.AddFunc(cronJob.GetCronExpression(), func() {
			// Execute the job handler and update LastRun
			err := cronJob.Handler(repos, params)
			if err != nil {
				log.Printf("Error executing job %s: %v", activeJob.JobName, err.Error())
			}

			updateJobStatus(repos, activeJob)
			log.Printf("Executed job: %s at %s\n", activeJob.JobName, time.Now())
		})

		if err != nil {
			log.Printf("Error scheduling job %s: %v\n", activeJob.JobName, err)
		}
	}

	// Start the cron scheduler
	c.Start()
	fmt.Println("Cron scheduler started.")
	return c
}

// Function to update last_ran and next_run
func updateJobStatus(repos *repository.Repositories, job model.Scheduler) {

	// Parse the cron expression
	// Join the last five fields for standard parsing (minute, hour, day of month, month, day of week)
	fields := strings.Fields(job.CronExpression)
	standardExpression := strings.Join(fields[1:], " ")
	schedule, err := cron.ParseStandard(standardExpression)
	if err != nil {
		log.Printf("Error parsing cron expression %s: %v\n", job.CronExpression, err)
		os.Exit(1)
	}

	now := time.Now()

	// Calculate next_run based on cron expression (example logic)
	nextRun := schedule.Next(time.Now()) // Adjust this based on your cron logic

	// Create an instance of the Scheduler model
	job.LastRun = &now
	job.NextRun = &nextRun

	// Update the job status in the database
	if err := repos.DatabaseRepos.Scheduler.Update(job); err != nil {
		log.Printf("Error updating job status for job ID %d: %v\n", job.ID, err)
	}
}
