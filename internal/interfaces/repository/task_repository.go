package repository

import "go-task-manager/internal/domain"

type TaskRepository interface {
	Create(task domain.Task) error
	GetAll() ([]domain.Task, error)
	GetByID(id uint) (domain.Task, error)
	Update(task domain.Task) error
	Delete(id uint) error
}
