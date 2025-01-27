package usecase

import (
	"errors"
	"go-task-manager/internal/domain"
	"go-task-manager/internal/interfaces/repository"
)

type taskUseCase struct {
	TaskRepo repository.TaskRepository
}

func NewTaskUseCase(taskRepo repository.TaskRepository) TaskUseCase {
	return &taskUseCase{
		TaskRepo: taskRepo,
	}
}

func (u *taskUseCase) CreateTask(task domain.Task) error {
	if task.Title == "" {
		return errors.New("ErrInvalidInput")
	}

	return u.TaskRepo.Create(task)
}

func (u *taskUseCase) GetTasks() ([]domain.Task, error) {
	return u.TaskRepo.GetAll()
}

func (u *taskUseCase) GetTaskByID(id uint) (domain.Task, error) {
	return u.TaskRepo.GetByID(id)
}

func (u *taskUseCase) UpdateTask(task domain.Task) error {
	return u.TaskRepo.Update(task)
}

func (u *taskUseCase) DeleteTask(id uint) error {
	return u.TaskRepo.Delete(id)
}
