package pgdb

import (
	"go-task-manager/internal/entity"
)

type TaskRepository interface {
	Create(task *entity.Task) error
	GetAll() ([]entity.Task, error)
	GetById(id uint) (entity.Task, error)
	Update(task entity.Task) error
	Delete(id uint) error
}
