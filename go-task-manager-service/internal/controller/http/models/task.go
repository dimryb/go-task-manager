package models

import (
	"go-task-manager/internal/entity"
	"time"
)

type CreateTaskRequest struct {
	Title       string    `json:"title" binding:"required" example:"Title"`
	Description string    `json:"description" example:"Description"`
	Status      string    `json:"status" binding:"required,oneof=pending in_progress done" example:"pending"`
	Priority    string    `json:"priority" binding:"required,oneof=low medium high" example:"medium"`
	DueDate     time.Time `json:"due_date" binding:"required" example:"2025-01-28T12:00:00Z"`
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
	taskResponses := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		taskResponses[i] = NewTaskResponse(task)
	}
	return taskResponses
}
