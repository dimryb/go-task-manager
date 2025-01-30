package tests

import (
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/assert"
	"go-task-manager-service/internal/repo/pgdb/models"
	"go-task-manager-service/tests"
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

	token := GetAuthToken(t)
	assert.NotEmpty(t, token)

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
		Header("Authorization", "Bearer "+token).
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

	token := GetAuthToken(t)
	assert.NotEmpty(t, token)

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
		Header("Authorization", "Bearer "+token).
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

	token := GetAuthToken(t)
	assert.NotEmpty(t, token)

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
		Header("Authorization", "Bearer "+token).
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

func TestGetFilteredTask_DateSuccess(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	db := app.DB

	token := GetAuthToken(t)
	assert.NotEmpty(t, token)

	dateStr := "2025-02-28T12:00:00Z"
	parsedTime, err := time.Parse(time.RFC3339, dateStr)
	assert.Nil(t, err)

	tasks := testTasks()
	tasks = append(tasks,
		models.Task{
			Title:       "Task Date",
			Description: "Date test task",
			Status:      "in_progress",
			Priority:    "medium",
			DueDate:     parsedTime,
		},
	)

	for _, task := range tasks {
		err := db.Create(&task).Error
		assert.Nil(t, err)
	}

	url := "/tasks"
	queryParams := map[string]string{
		"due_date": dateStr,
	}

	apitest.
		Handler(app.Router).
		Get(url).
		Header("Authorization", "Bearer "+token).
		QueryParams(queryParams).
		Expect(t).
		Status(http.StatusOK).
		Header("Content-Type", "application/json").
		Assert(jsonpath.Len("result", 1)).
		Assert(
			jsonpath.Chain().
				Equal("ok", true).
				Equal("result[0].title", "Task Date").
				Equal("result[0].description", "Date test task").
				Equal("result[0].status", "in_progress").
				Equal("result[0].priority", "medium").
				Equal("result[0].due_date", dateStr).
				End(),
		).
		End()
}
