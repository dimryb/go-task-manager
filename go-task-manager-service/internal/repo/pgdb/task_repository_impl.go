package pgdb

import (
	"go-task-manager/internal/entity"
	"go-task-manager/internal/repo/pgdb/models"
	"gorm.io/gorm"
)

type taskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{DB: db}
}

func (r *taskRepository) Create(task *entity.Task) error {
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

func (r *taskRepository) GetAll() ([]entity.Task, error) {
	return make([]entity.Task, 0), nil //todo: реализовать
}

func (r *taskRepository) GetById(id uint) (entity.Task, error) {
	var dbTask models.Task
	err := r.DB.First(&dbTask, id).Error
	if err != nil {
		return entity.Task{}, err
	}

	task := entity.Task{
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

func (r *taskRepository) Update(task entity.Task) error {
	return nil //todo: реализовать
}

func (r *taskRepository) Delete(id uint) error {
	return nil //todo: реализовать
}
