package repository

import (
	"github.com/admiralyeoj/animanager/internal/database/model"
	"gorm.io/gorm"
)

type SchedulerRepository interface {
	// Define your methods here
	GetAll() ([]model.Scheduler, error)
	Update(schedule model.Scheduler) error
}

// schedulerRepository is a concrete implementation of SchedulerRepository
type schedulerRepository struct {
	db *gorm.DB
}

func NewSchedulerRepository(db *gorm.DB) SchedulerRepository {
	return &schedulerRepository{
		db: db,
	}
}

func (schedulerRepo *schedulerRepository) GetAll() ([]model.Scheduler, error) {
	var jobs []model.Scheduler
	// Fetch active jobs using GORM's `Where` and `Find` methods
	if err := schedulerRepo.db.Where("is_active = ?", true).Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

func (schedulerRepo *schedulerRepository) Update(schedule model.Scheduler) error {
	return schedulerRepo.db.Model(&schedule).Updates(schedule).Error
}
