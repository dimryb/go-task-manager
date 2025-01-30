package service

import (
	"errors"
	"go-task-manager/internal/entity"
	"go-task-manager/internal/repo/pgdb"
)

type taskUseCase struct {
	TaskRepo pgdb.TaskRepository
}

func NewTaskUseCase(taskRepo pgdb.TaskRepository) TaskUseCase {
	return &taskUseCase{
		TaskRepo: taskRepo,
	}
}

func (u *taskUseCase) CreateTask(task *entity.Task) error {
	if err := u.TaskRepo.Create(task); err != nil {
		return errors.New("failed to create task: " + err.Error())
	}
	return nil
}

func (u *taskUseCase) GetTasksFiltered(status, priority, dueDate, title string) ([]entity.Task, error) {
	tasks, err := u.TaskRepo.GetFiltered(status, priority, dueDate, title)
	if err != nil {
		return nil, errors.New("failed to get tasks filtered: " + err.Error())
	}
	return tasks, nil
}

func (u *taskUseCase) GetTaskByID(id uint) (entity.Task, error) {
	return u.TaskRepo.GetById(id)
}

func (u *taskUseCase) UpdateTask(task entity.Task) error {
	if err := u.TaskRepo.Update(task); err != nil {
		return errors.New("failed to update task: " + err.Error())
	}
	return nil
}

func (u *taskUseCase) DeleteTask(id uint) error {
	return u.TaskRepo.Delete(id)
}

func (u *taskUseCase) GetAllTasks() ([]entity.Task, error) {
	tasks, err := u.TaskRepo.GetFiltered("", "", "", "")
	if err != nil {
		return nil, errors.New("failed to get all tasks: " + err.Error())
	}
	return tasks, nil
}

func (u *taskUseCase) CreateTasks(tasks []entity.Task) error {
	for _, task := range tasks {
		if err := u.CreateTask(&task); err != nil {
			return errors.New("failed to create tasks: " + err.Error())
		}
	}
	return nil
}
