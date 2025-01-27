package handler

import (
	"go-task-manager/internal/usecase"
	"net/http"
)

type TaskHandler struct {
	TaskUseCase usecase.TaskUseCase
}

func (h TaskHandler) CreateTask(writer http.ResponseWriter, request *http.Request) {
	//todo: реализовать
}
