package tests

import (
	"fmt"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
	"go-task-manager-service/internal/repo/pgdb/models"
	"go-task-manager-service/tests"
	"net/http"
	"testing"
	"time"
)

func TestDeleteTask_Success(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	db := app.DB

	token := GetAuthToken(t)
	assert.NotEmpty(t, token)

	task := models.Task{
		Title:       "Task to Delete",
		Description: "Will be deleted",
		Status:      "pending",
		Priority:    "medium",
		DueDate:     time.Now(),
	}
	err := db.Create(&task).Error
	assert.Nil(t, err)

	url := fmt.Sprintf("/tasks/%d", task.ID)

	apitest.
		Handler(app.Router).
		Delete(url).
		Header("Authorization", "Bearer "+token).
		Expect(t).
		Status(http.StatusOK).
		Body(`{"ok": true, "result": 1}`).
		End()
}

func TestDeleteTask_NotFound(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	token := GetAuthToken(t)
	assert.NotEmpty(t, token)

	url := "/tasks/9999"

	apitest.
		Handler(app.Router).
		Delete(url).
		Header("Authorization", "Bearer "+token).
		Expect(t).
		Status(http.StatusNotFound).
		Body(`{"ok": false, "error": "task not found"}`).
		End()
}

func TestDeleteTask_InvalidID(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	token := GetAuthToken(t)
	assert.NotEmpty(t, token)

	url := "/tasks/abc" // Некорректный ID

	apitest.
		Handler(app.Router).
		Delete(url).
		Header("Authorization", "Bearer "+token).
		Expect(t).
		Status(http.StatusBadRequest).
		Body(`{"ok": false, "error": "invalid task ID"}`).
		End()
}
