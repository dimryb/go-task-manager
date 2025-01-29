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

// CreateTask godoc
// @Summary Создание задачи
// @Description Создает новую задачу
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body models.CreateTaskRequest true "Создание задачи"
// @Success 201 {object} rest.Response
// @Failure 400 {object} rest.Response "Некорректный запрос"
// @Failure 500 {object} rest.Response "Ошибка сервера"
// @Router /tasks [post]
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

// GetTaskById godoc
// @Summary Получение задачи по ID (вспомогательная)
// @Description Возвращает задачу по её ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} rest.Response
// @Failure 400 {object} rest.Response "Некорректный ID"
// @Failure 404 {object} rest.Response "Задача не найдена"
// @Failure 500 {object} rest.Response "Ошибка сервера"
// @Router /tasks/{id} [get]
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

// GetTasksFiltered godoc
// @Summary Получение списка задач с фильтрацией
// @Description Возвращает задачи, отфильтрованные по статусу, приоритету и дате
// @Tags tasks
// @Accept json
// @Produce json
// @Param status query string false "Статус (pending, in_progress, done)"
// @Param priority query string false "Приоритет (low, medium, high)"
// @Param due_date query string false "Дата выполнения формат: 2020-01-01T12:00:00Z"
// @Param title query string false "Название задачи"
// @Success 200 {object} rest.Response
// @Failure 400 {object} rest.Response "Некорректные параметры запроса"
// @Failure 500 {object} rest.Response "Ошибка сервера"
// @Router /tasks [get]
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

// UpdateTask godoc
// @Summary Обновление задачи
// @Description Обновляет информацию о задаче
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param task body models.UpdateTaskRequest true "Данные для обновления"
// @Success 200 {object} rest.Response
// @Failure 400 {object} rest.Response "Некорректные данные"
// @Failure 404 {object} rest.Response "Задача не найдена"
// @Failure 500 {object} rest.Response "Ошибка сервера"
// @Router /tasks/{id} [put]
func (h taskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		rest.WriteError(w, http.StatusBadRequest, errors.New("invalid task ID"))
		return
	}

	var req models.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rest.WriteError(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	if err := validateUpdateTaskRequest(req); err != nil {
		rest.WriteError(w, http.StatusBadRequest, err)
		return
	}

	task, err := h.TaskUseCase.GetTaskByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rest.WriteError(w, http.StatusNotFound, errors.New("task not found"))
			return
		}
		rest.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Status != nil {
		task.Status = *req.Status
	}
	if req.Priority != nil {
		task.Priority = *req.Priority
	}
	if req.DueDate != nil {
		task.DueDate = *req.DueDate
	}

	if err := h.TaskUseCase.UpdateTask(task); err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func validateUpdateTaskRequest(req models.UpdateTaskRequest) error {
	if req.Title != nil && *req.Title == "" {
		return errors.New("title is required")
	}
	if req.Status != nil && *req.Status != "pending" && *req.Status != "in_progress" && *req.Status != "done" {
		return errors.New("status must be one of: pending, in_progress, done")
	}
	if req.Priority != nil && *req.Priority != "low" && *req.Priority != "medium" && *req.Priority != "high" {
		return errors.New("priority must be one of: low, medium, high")
	}

	return nil
}

func (h taskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
