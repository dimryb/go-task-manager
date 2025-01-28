package tests

import (
	"go-task-manager/internal/domain"
	"go-task-manager/internal/infrastructure/http/handler"
	"go-task-manager/internal/infrastructure/http/models"
	"net/http"
	"testing"
	"time"

	"github.com/steinfletcher/apitest"
)

func TestCreateTask_Success(t *testing.T) {
	mockUseCase := &mockTaskUseCase{
		CreateTaskFn: func(task domain.Task) error {
			if task.Title != "Test Task" || task.Status != "pending" || task.Priority != "high" {
				t.Fatalf("Task data is incorrect")
			}
			return nil
		},
	}

	taskHandler := handler.NewTaskHandler(mockUseCase)
	reqBody := models.CreateTaskRequest{
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      "pending",
		Priority:    "high",
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	apitest.New().
		HandlerFunc(taskHandler.CreateTask).
		Post("/tasks").
		JSON(reqBody).
		Expect(t).
		Status(http.StatusCreated).
		Body("Task created successfully").
		End()
}

func TestCreateTask_InvalidInput(t *testing.T) {
	mockUseCase := &mockTaskUseCase{}
	taskHandler := handler.NewTaskHandler(mockUseCase)

	apitest.New().
		HandlerFunc(taskHandler.CreateTask).
		Post("/tasks").
		JSON(`{ invalid json }`).
		Expect(t).
		Status(http.StatusBadRequest).
		Body("Invalid input\n").
		End()
}

func TestCreateTask_ValidationTitleEmpty(t *testing.T) {
	mockUseCase := &mockTaskUseCase{}

	taskHandler := handler.NewTaskHandler(mockUseCase)
	reqBody := models.CreateTaskRequest{
		Title:    "",
		Status:   "invalid_status",
		Priority: "low",
		DueDate:  time.Now().Add(24 * time.Hour),
	}

	apitest.New().
		HandlerFunc(taskHandler.CreateTask).
		Post("/tasks").
		JSON(reqBody).
		Expect(t).
		Status(http.StatusBadRequest).
		Body("title is required\n").
		End()
}

func TestCreateTask_ValidationInvalidStatus(t *testing.T) {
	mockUseCase := &mockTaskUseCase{}

	taskHandler := handler.NewTaskHandler(mockUseCase)
	reqBody := models.CreateTaskRequest{
		Title:    "Test Task",
		Status:   "invalid_status",
		Priority: "low",
		DueDate:  time.Now().Add(24 * time.Hour),
	}

	apitest.New().
		HandlerFunc(taskHandler.CreateTask).
		Post("/tasks").
		JSON(reqBody).
		Expect(t).
		Status(http.StatusBadRequest).
		Body("status must be one of: pending, in_progress, done\n").
		End()
}

func TestCreateTask_ValidationInvalidPriority(t *testing.T) {
	mockUseCase := &mockTaskUseCase{}

	taskHandler := handler.NewTaskHandler(mockUseCase)
	reqBody := models.CreateTaskRequest{
		Title:    "Test Task",
		Status:   "pending",
		Priority: "invalid_priority",
		DueDate:  time.Now().Add(24 * time.Hour),
	}

	apitest.New().
		HandlerFunc(taskHandler.CreateTask).
		Post("/tasks").
		JSON(reqBody).
		Expect(t).
		Status(http.StatusBadRequest).
		Body("priority must be one of: low, medium, high\n").
		End()
}
