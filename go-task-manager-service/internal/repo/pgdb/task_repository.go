package pgdb

import (
	"go-task-manager-service/internal/entity"
)

type TaskRepository interface {
	Create(task *entity.Task) error
	GetFiltered(status, priority, dueDate, title string) ([]entity.Task, error)
	GetById(id uint) (entity.Task, error)
	Update(task entity.Task) error
	Delete(id uint) error
}
