package usecase

import (
	"go-task-manager/internal/domain"
)

type TaskUseCase interface {
	CreateTask(task *domain.Task) error
	GetTasks() ([]domain.Task, error)
	GetTaskByID(id uint) (domain.Task, error)
	UpdateTask(task domain.Task) error
	DeleteTask(id uint) error
}
