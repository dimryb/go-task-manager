package repository

import (
	"go-task-manager/internal/domain"
	"go-task-manager/internal/interfaces/database/models"
	"gorm.io/gorm"
)

type taskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{DB: db}
}

func (r *taskRepository) Create(task domain.Task) error {
	dbTask := models.Task{
		Title:    task.Title,
		Status:   task.Status,
		Priority: task.Priority,
		DueDate:  task.DueDate,
	}
	return r.DB.Create(&dbTask).Error
}

func (r *taskRepository) GetAll() ([]domain.Task, error) {
	return make([]domain.Task, 0), nil //todo: реализовать
}

func (r *taskRepository) GetByID(id uint) (domain.Task, error) {
	return domain.Task{}, nil //todo: реализовать
}

func (r *taskRepository) Update(task domain.Task) error {
	return nil //todo: реализовать
}

func (r *taskRepository) Delete(id uint) error {
	return nil //todo: реализовать
}
