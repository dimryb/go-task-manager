package handler

import (
	"encoding/json"
	"go-task-manager/internal/domain"
	"go-task-manager/internal/interfaces/http/models"
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

func (h taskHandler) CreateTask(writer http.ResponseWriter, request *http.Request) {
	var req models.CreateTaskRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, "Invalid input", http.StatusBadRequest)
		return
	}

	task := domain.Task{}

	if err := h.TaskUseCase.CreateTask(task); err != nil {
		http.Error(writer, "Failed to create task", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}
