package service

import (
	"go-task-manager/internal/entity"
)

type TaskUseCase interface {
	CreateTask(task *entity.Task) error
	GetTasks() ([]entity.Task, error)
	GetTaskByID(id uint) (entity.Task, error)
	UpdateTask(task entity.Task) error
	DeleteTask(id uint) error
}
