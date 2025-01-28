package tests

import (
	"go-task-manager/internal/domain"
)

type mockTaskUseCase struct {
	CreateTaskFn  func(task *domain.Task) error
	GetTasksFn    func() ([]domain.Task, error)
	GetTaskByIDFn func(id uint) (domain.Task, error)
	UpdateTaskFn  func(task domain.Task) error
	DeleteTaskFn  func(id uint) error
}

func (m *mockTaskUseCase) CreateTask(task *domain.Task) error {
	if m.CreateTaskFn != nil {
		return m.CreateTaskFn(task)
	}
	return nil
}

func (m *mockTaskUseCase) GetTasks() ([]domain.Task, error) {
	if m.GetTasksFn != nil {
		return m.GetTasksFn()
	}
	return nil, nil
}

func (m *mockTaskUseCase) GetTaskByID(id uint) (domain.Task, error) {
	if m.GetTaskByIDFn != nil {
		return m.GetTaskByIDFn(id)
	}
	return domain.Task{}, nil
}

func (m *mockTaskUseCase) UpdateTask(task domain.Task) error {
	if m.UpdateTaskFn != nil {
		return m.UpdateTaskFn(task)
	}
	return nil
}

func (m *mockTaskUseCase) DeleteTask(id uint) error {
	if m.DeleteTaskFn != nil {
		return m.DeleteTaskFn(id)
	}
	return nil
}
