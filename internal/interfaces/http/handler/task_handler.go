package handler

import (
	"encoding/json"
	"go-task-manager/internal/domain"
	"go-task-manager/internal/usecase"
	"net/http"
)

type TaskHandler struct {
	TaskUseCase usecase.TaskUseCase
}

func (h TaskHandler) CreateTask(writer http.ResponseWriter, request *http.Request) {
	var req domain.Task
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		http.Error(writer, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.TaskUseCase.CreateTask(req); err != nil {
		http.Error(writer, "Failed to create task", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}
