package tests

import (
	"go-task-manager/internal/entity"
)

type mockTaskUseCase struct {
	CreateTaskFn  func(task *entity.Task) error
	GetTasksFn    func() ([]entity.Task, error)
	GetTaskByIDFn func(id uint) (entity.Task, error)
	UpdateTaskFn  func(task entity.Task) error
	DeleteTaskFn  func(id uint) error
}

func (m *mockTaskUseCase) CreateTask(task *entity.Task) error {
	if m.CreateTaskFn != nil {
		return m.CreateTaskFn(task)
	}
	return nil
}

func (m *mockTaskUseCase) GetTasks() ([]entity.Task, error) {
	if m.GetTasksFn != nil {
		return m.GetTasksFn()
	}
	return nil, nil
}

func (m *mockTaskUseCase) GetTaskByID(id uint) (entity.Task, error) {
	if m.GetTaskByIDFn != nil {
		return m.GetTaskByIDFn(id)
	}
	return entity.Task{}, nil
}

func (m *mockTaskUseCase) UpdateTask(task entity.Task) error {
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
