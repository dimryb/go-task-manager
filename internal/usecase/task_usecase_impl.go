package usecase

import (
	"errors"
	"go-task-manager/internal/domain"
	"go-task-manager/internal/infrastructure/repository"
)

type taskUseCase struct {
	TaskRepo repository.TaskRepository
}

func NewTaskUseCase(taskRepo repository.TaskRepository) TaskUseCase {
	return &taskUseCase{
		TaskRepo: taskRepo,
	}
}

func (u *taskUseCase) CreateTask(task *domain.Task) error {
	if err := u.TaskRepo.Create(task); err != nil {
		return errors.New("failed to create task: " + err.Error())
	}
	return nil
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
