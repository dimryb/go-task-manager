package repository

import (
	"go-task-manager/internal/domain"
	"go-task-manager/internal/infrastructure/database/models"
	"gorm.io/gorm"
)

type taskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{DB: db}
}

func (r *taskRepository) Create(task *domain.Task) error {
	dbTask := models.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		DueDate:     task.DueDate,
	}
	err := r.DB.Create(&dbTask).Error
	if err == nil {
		task.ID = dbTask.ID
		task.CreatedAt = dbTask.CreatedAt
		task.UpdatedAt = dbTask.UpdatedAt
	}
	return err
}

func (r *taskRepository) GetAll() ([]domain.Task, error) {
	return make([]domain.Task, 0), nil //todo: реализовать
}

func (r *taskRepository) GetById(id uint) (domain.Task, error) {
	var dbTask models.Task
	err := r.DB.First(&dbTask, id).Error
	if err != nil {
		return domain.Task{}, err
	}

	task := domain.Task{
		ID:          dbTask.ID,
		Title:       dbTask.Title,
		Description: dbTask.Description,
		Status:      dbTask.Status,
		Priority:    dbTask.Priority,
		DueDate:     dbTask.DueDate,
		CreatedAt:   dbTask.CreatedAt,
		UpdatedAt:   dbTask.UpdatedAt,
	}
	return task, nil
}

func (r *taskRepository) Update(task domain.Task) error {
	return nil //todo: реализовать
}

func (r *taskRepository) Delete(id uint) error {
	return nil //todo: реализовать
}
