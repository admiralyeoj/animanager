package cron

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/admiralyeoj/animanager/internal/config"
	"github.com/admiralyeoj/animanager/internal/repository"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

func StartCron(cfg *config.Config, db *gorm.DB) {

	// Create a root command to serve as the entry point
	repos := repository.InitializeRepositories(cfg, db)
	// srvs := service.InitializeServices(repos, db)

	// Load jobs from the scheduler table
	jobs, err := repos.DatabaseRepos.Scheduler.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	if len(jobs) == 0 {
		fmt.Println("No Jobs Found.")
		return
	}

	// Map function identifiers to actual functions
	funcMap := map[string]func(map[string]interface{}){
		"backupDatabase":    backupDatabase,
		"sendNotifications": sendNotifications,
	}

	// Create a new cron instance
	c := cron.New()

	// Schedule each job based on the cron expression and function mapping
	for _, job := range jobs {
		job := job // capture job variable for closure

		// Get the function from the map based on `function_name`
		jobFunc, exists := funcMap[job.FunctionName]
		if !exists {
			log.Printf("No function found for job %s, skipping...\n", job.JobName)
			continue
		}

		fmt.Println(job.JobName)

		// Add the job to the cron scheduler
		fmt.Println(job.CronExpression)
		err := c.AddFunc(job.CronExpression, func() {

			var params map[string]interface{}

			// Unmarshal params from JSON
			if err := json.Unmarshal(job.Params, &params); err != nil {
				log.Printf("Error unmarshalling params for job %s: %v\n", job.JobName, err)
				return // Early exit on error
			}

			jobFunc(params)
			fmt.Printf("Executed job: %s\n", job.JobName)
		})

		if err != nil {
			log.Printf("Error scheduling job %s: %v\n", job.JobName, err)
		}
	}

	// Start the cron scheduler
	c.Start()
	fmt.Println("Starting cron")

	// Keep the application running
	select {}
}

func backupDatabase(params map[string]interface{}) {
	fmt.Println("Running database backup with params:", params)
	// Actual backup logic here
}

func sendNotifications(params map[string]interface{}) {
	fmt.Println("Sending notifications with params:", params)
	// Actual notification logic here
}
