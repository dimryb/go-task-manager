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
		Result: models.NewTaskResponse(task),
	})
}

func (h taskHandler) GetTasksFiltered(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	status := queryParams.Get("status")
	priority := queryParams.Get("priority")
	dueDate := queryParams.Get("due_date")
	title := queryParams.Get("title")

	if err := validateGetTasksFiltered(status, priority, dueDate); err != nil {
		rest.WriteError(w, http.StatusBadRequest, err)
		return
	}

	tasks, err := h.TaskUseCase.GetTasksFiltered(status, priority, dueDate, title)
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	rest.WriteJSON(w, http.StatusOK, rest.Response{
		Ok:     true,
		Result: models.NewTasksResponse(tasks),
	})
}

func validateGetTasksFiltered(status, priority, dueDate string) error {
	if status != "pending" && status != "in_progress" && status != "done" && status != "" {
		return errors.New("status must be one of: pending, in_progress, done")
	}
	if priority != "low" && priority != "medium" && priority != "high" && priority != "" {
		return errors.New("priority must be one of: low, medium, high")
	}

	return nil
}

func (h taskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h taskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
