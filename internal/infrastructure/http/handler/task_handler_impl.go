package handler

import (
	"encoding/json"
	"errors"
	"go-task-manager/internal/domain"
	"go-task-manager/internal/infrastructure/http/models"
	"go-task-manager/internal/usecase"
	"net/http"
)

type taskHandler struct {
	TaskUseCase usecase.TaskUseCase
}

func NewTaskHandler(useCase usecase.TaskUseCase) TaskHandler {
	return &taskHandler{
		TaskUseCase: useCase,
	}
}

func (h taskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := validateCreateTaskRequest(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := domain.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
	}

	if err := h.TaskUseCase.CreateTask(task); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task created successfully"))
}

func validateCreateTaskRequest(req models.CreateTaskRequest) error {
	if req.Title == "" {
		return errors.New("title is required")
	}
	if req.Status != "pending" && req.Status != "in_progress" && req.Status != "done" {
		return errors.New("status must be one of: pending, in_progress, done")
	}
	if req.Priority != "low" && req.Priority != "medium" && req.Priority != "high" {
		return errors.New("priority must be one of: low, medium, high")
	}

	return nil
}
