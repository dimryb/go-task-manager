package repository

import (
	"go-task-manager/internal/domain"
	"go-task-manager/internal/interfaces/database/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func (r *TaskRepository) Create(task domain.Task) error {
	dbTask := models.Task{
		Title:    task.Title,
		Status:   task.Status,
		Priority: task.Priority,
		DueDate:  task.DueDate,
	}
	return r.DB.Create(&dbTask).Error
}

func (r *TaskRepository) GetAll() ([]domain.Task, error) {
	return make([]domain.Task, 0), nil //todo: реализовать
}

func (r *TaskRepository) GetByID(id uint) (domain.Task, error) {
	return domain.Task{}, nil //todo: реализовать
}

func (r *TaskRepository) Update(task domain.Task) error {
	return nil //todo: реализовать
}

func (r *TaskRepository) Delete(id uint) error {
	return nil //todo: реализовать
}
