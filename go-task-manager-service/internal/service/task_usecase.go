package service

import (
	"go-task-manager-service/internal/entity"
)

type TaskUseCase interface {
	CreateTask(task *entity.Task) error
	GetTasksFiltered(status, priority, dueDate, title string) ([]entity.Task, error)
	GetTaskByID(id uint) (entity.Task, error)
	UpdateTask(task entity.Task) error
	DeleteTask(id uint) error
	GetAllTasks() ([]entity.Task, error)
	CreateTasks(tasks []entity.Task) error
}
