package cron

import (
	"github.com/admiralyeoj/animanager/internal/config"
	"github.com/admiralyeoj/animanager/internal/cronJob"
	aniRepo "github.com/admiralyeoj/animanager/internal/repository"
	"gorm.io/gorm"
)

func StartCron(cfg *config.Config, db *gorm.DB) {

	// Create a root command to serve as the entry point
	repos := aniRepo.InitializeRepositories(cfg, db)
	// srvs := service.InitializeServices(repos, db)

	cronJob.InitializeCronJobs(repos)

	// Keep the application running
	select {}
}
