package models

import (
	"go-task-manager-service/internal/entity"
	"go-task-manager-service/pkg/utils"
	"time"
)

type CreateTaskRequest struct {
	Title       string    `json:"title" binding:"required" example:"Title"`
	Description string    `json:"description" example:"Description"`
	Status      string    `json:"status" binding:"required,oneof=pending in_progress done" example:"pending"`
	Priority    string    `json:"priority" binding:"required,oneof=low medium high" example:"medium"`
	DueDate     time.Time `json:"due_date" binding:"required" example:"2025-01-28T12:00:00Z"`
}

type UpdateTaskRequest struct {
	Title       *string    `json:"title" example:"Updated Title"`
	Description *string    `json:"description" example:"Updated Description"`
	Status      *string    `json:"status" binding:"oneof=pending in_progress done" example:"in_progress"`
	Priority    *string    `json:"priority" binding:"oneof=low medium high" example:"high"`
	DueDate     *time.Time `json:"due_date" example:"2025-02-01T15:00:00Z"`
}

type TaskResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	DueDate     string `json:"due_date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateTaskImportRequest struct {
	ID          uint   `json:"id" binding:"required" example:"1"`
	Title       string `json:"title" binding:"required" example:"Title"`
	Description string `json:"description" example:"Description"`
	Status      string `json:"status" binding:"required,oneof=pending in_progress done" example:"pending"`
	Priority    string `json:"priority" binding:"required,oneof=low medium high" example:"medium"`
	DueDate     utils.JSONTime
	CreatedAt   utils.JSONTime
	UpdatedAt   utils.JSONTime
}

func NewTaskResponse(task entity.Task) TaskResponse {
	return TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		DueDate:     task.DueDate.Format(time.RFC3339),
		CreatedAt:   task.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
	}
}

func NewTasksResponse(tasks []entity.Task) []TaskResponse {
	tasksResponses := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		tasksResponses[i] = NewTaskResponse(task)
	}
	return tasksResponses
}

func NewTaskEntity(task CreateTaskImportRequest) entity.Task {
	return entity.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		DueDate:     task.DueDate.Time(),
		CreatedAt:   task.CreatedAt.Time(),
		UpdatedAt:   task.UpdatedAt.Time(),
	}
}

func NewTasksEntity(tasks []CreateTaskImportRequest) []entity.Task {
	tasksEntity := make([]entity.Task, len(tasks))
	for i, task := range tasks {
		tasksEntity[i] = NewTaskEntity(task)
	}
	return tasksEntity
}
