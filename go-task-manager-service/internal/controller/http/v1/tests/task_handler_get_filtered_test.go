package tests

import (
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/assert"
	"go-task-manager/internal/repo/pgdb/models"
	"go-task-manager/tests"
	"net/http"
	"testing"
	"time"
)

func testTasks() []models.Task {
	tasks := []models.Task{
		{
			Title:       "Task 1",
			Description: "First test task",
			Status:      "pending",
			Priority:    "low",
			DueDate:     time.Now(),
		},
		{
			Title:       "Task 2",
			Description: "Second test task",
			Status:      "done",
			Priority:    "high",
			DueDate:     time.Now(),
		},
		{
			Title:       "Task 3",
			Description: "Third test task",
			Status:      "done",
			Priority:    "high",
			DueDate:     time.Now(),
		},
	}
	return tasks
}

func TestGetFilteredTask_Success(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	db := app.DB

	tasks := testTasks()

	for _, task := range tasks {
		err := db.Create(&task).Error
		assert.Nil(t, err)
	}

	url := "/tasks"
	queryParams := map[string]string{
		"title":    "Task",
		"status":   "done",
		"priority": "high",
	}

	apitest.
		Handler(app.Router).
		Get(url).
		QueryParams(queryParams).
		Expect(t).
		Status(http.StatusOK).
		Header("Content-Type", "application/json").
		Assert(jsonpath.Len("result", 2)).
		Assert(
			jsonpath.Chain().
				Equal("ok", true).
				Equal("result[0].title", "Task 2").
				Equal("result[0].description", "Second test task").
				Equal("result[0].status", "done").
				Equal("result[0].priority", "high").
				Equal("result[1].title", "Task 3").
				Equal("result[1].description", "Third test task").
				Equal("result[1].status", "done").
				Equal("result[1].priority", "high").
				End(),
		).
		End()
}

func TestGetFilteredTask_ValidationInvalidStatus(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	db := app.DB

	tasks := testTasks()

	for _, task := range tasks {
		err := db.Create(&task).Error
		assert.Nil(t, err)
	}

	url := "/tasks"
	queryParams := map[string]string{
		"title":    "Task",
		"status":   "invalid_status",
		"priority": "high",
	}

	apitest.
		Handler(app.Router).
		Get(url).
		QueryParams(queryParams).
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

func TestGetFilteredTask_ValidationInvalidPriority(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	db := app.DB

	tasks := testTasks()

	for _, task := range tasks {
		err := db.Create(&task).Error
		assert.Nil(t, err)
	}

	url := "/tasks"
	queryParams := map[string]string{
		"title":    "Task",
		"status":   "done",
		"priority": "invalid_priority",
	}

	apitest.
		Handler(app.Router).
		Get(url).
		QueryParams(queryParams).
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
