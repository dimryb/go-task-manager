package tests

import (
	"go-task-manager/internal/controller/http/models"
	"go-task-manager/internal/controller/http/v1"
	"go-task-manager/internal/entity"
	"net/http"
	"testing"
	"time"

	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestCreateTask_Success(t *testing.T) {
	mockUseCase := &mockTaskUseCase{
		CreateTaskFn: func(task *entity.Task) error {
			if task.Title != "Test Task" || task.Status != "pending" || task.Priority != "high" {
				t.Fatalf("Task data is incorrect")
			}
			task.ID = 1
			return nil
		},
	}

	taskHandler := v1.NewTaskHandler(mockUseCase)
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
		Header("Content-Type", "application/json").
		Assert(
			jsonpath.Chain().
				Equal("ok", true).
				Equal("result", float64(1)).
				End(),
		).
		End()
}

func TestCreateTask_InvalidInput(t *testing.T) {
	mockUseCase := &mockTaskUseCase{}
	taskHandler := v1.NewTaskHandler(mockUseCase)

	apitest.New().
		HandlerFunc(taskHandler.CreateTask).
		Post("/tasks").
		JSON(`{ invalid json }`).
		Expect(t).
		Status(http.StatusBadRequest).
		Header("Content-Type", "application/json").
		Assert(
			jsonpath.Chain().
				Equal("ok", false).
				Present("error").
				End(),
		).
		End()
}

func TestCreateTask_ValidationTitleEmpty(t *testing.T) {
	mockUseCase := &mockTaskUseCase{}

	taskHandler := v1.NewTaskHandler(mockUseCase)
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
		Header("Content-Type", "application/json").
		Assert(
			jsonpath.Chain().
				Equal("ok", false).
				Equal("error", "title is required").
				End(),
		).
		End()
}

func TestCreateTask_ValidationInvalidStatus(t *testing.T) {
	mockUseCase := &mockTaskUseCase{}

	taskHandler := v1.NewTaskHandler(mockUseCase)
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
		Header("Content-Type", "application/json").
		Assert(
			jsonpath.Chain().
				Equal("ok", false).
				Equal("error", "status must be one of: pending, in_progress, done").
				End(),
		).
		End()
}

func TestCreateTask_ValidationInvalidPriority(t *testing.T) {
	mockUseCase := &mockTaskUseCase{}

	taskHandler := v1.NewTaskHandler(mockUseCase)
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
		Header("Content-Type", "application/json").
		Assert(
			jsonpath.Chain().
				Equal("ok", false).
				Equal("error", "priority must be one of: low, medium, high").
				End(),
		).
		End()
}
