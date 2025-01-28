package v1

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"go-task-manager/internal/controller/http/models"
	"go-task-manager/internal/controller/http/rest"
	"go-task-manager/internal/entity"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"go-task-manager/internal/service"
)

type taskHandler struct {
	TaskUseCase service.TaskUseCase
}

func NewTaskHandler(useCase service.TaskUseCase) TaskHandler {
	return &taskHandler{
		TaskUseCase: useCase,
	}
}

func (h taskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rest.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := validateCreateTaskRequest(req); err != nil {
		rest.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task := entity.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
	}

	if err := h.TaskUseCase.CreateTask(&task); err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	rest.WriteJSON(w, http.StatusCreated, rest.Response{
		Ok:     true,
		Result: task.ID,
	})
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

func (h taskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.TaskUseCase.GetTaskByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rest.WriteError(w, http.StatusNotFound, errors.New("not found"))
			return
		}
		rest.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	rest.WriteJSON(w, http.StatusOK, rest.Response{
		Ok:     true,
		Result: task, // todo: в response перевести
	})
}
