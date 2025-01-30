package tests

import (
	"go-task-manager-service/internal/entity"
)

type MockTaskUseCase struct {
	CreateTaskFn       func(task *entity.Task) error
	GetTasksFilteredFn func(status, priority, dueDate, title string) ([]entity.Task, error)
	GetTaskByIDFn      func(id uint) (entity.Task, error)
	UpdateTaskFn       func(task entity.Task) error
	DeleteTaskFn       func(id uint) error
}

func (m *MockTaskUseCase) CreateTask(task *entity.Task) error {
	if m.CreateTaskFn != nil {
		return m.CreateTaskFn(task)
	}
	return nil
}

func (m *MockTaskUseCase) GetTasksFiltered(status, priority, dueDate, title string) ([]entity.Task, error) {
	if m.GetTasksFilteredFn != nil {
		return m.GetTasksFilteredFn(status, priority, dueDate, title)
	}
	return nil, nil
}

func (m *MockTaskUseCase) GetTaskByID(id uint) (entity.Task, error) {
	if m.GetTaskByIDFn != nil {
		return m.GetTaskByIDFn(id)
	}
	return entity.Task{}, nil
}

func (m *MockTaskUseCase) UpdateTask(task entity.Task) error {
	if m.UpdateTaskFn != nil {
		return m.UpdateTaskFn(task)
	}
	return nil
}

func (m *MockTaskUseCase) DeleteTask(id uint) error {
	if m.DeleteTaskFn != nil {
		return m.DeleteTaskFn(id)
	}
	return nil
}
