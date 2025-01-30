package tests

import (
	"fmt"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
	pgModels "go-task-manager-service/internal/repo/pgdb/models"
	"go-task-manager-service/tests"
	"net/http"
	"testing"
	"time"
)

func TestUpdateTask_Success(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	db := app.DB

	token := GetAuthToken(t)
	assert.NotEmpty(t, token)

	task := pgModels.Task{
		Title:       "Old Title",
		Description: "Old Description",
		Status:      "pending",
		Priority:    "medium",
		DueDate:     time.Now(),
	}
	err := db.Create(&task).Error
	assert.Nil(t, err)

	url := fmt.Sprintf("/tasks/%d", task.ID)

	updateRequest := map[string]interface{}{
		"title":       "Updated Title",
		"description": "Updated Description",
		"status":      "done",
		"priority":    "high",
	}

	apitest.
		Handler(app.Router).
		Put(url).
		Header("Authorization", "Bearer "+token).
		JSON(updateRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdateTask_NotFound(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	token := GetAuthToken(t)
	assert.NotEmpty(t, token)

	url := "/tasks/9999"

	updateRequest := map[string]interface{}{
		"title": "Updated Title",
	}

	apitest.
		Handler(app.Router).
		Put(url).
		Header("Authorization", "Bearer "+token).
		JSON(updateRequest).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestUpdateTask_InvalidID(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	token := GetAuthToken(t)
	assert.NotEmpty(t, token)

	url := "/tasks/abc"

	updateRequest := map[string]interface{}{
		"title": "Updated Title",
	}

	apitest.
		Handler(app.Router).
		Put(url).
		Header("Authorization", "Bearer "+token).
		JSON(updateRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdateTask_InvalidData(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	token := GetAuthToken(t)
	assert.NotEmpty(t, token)

	db := app.DB

	task := pgModels.Task{
		Title:       "Old Title",
		Description: "Old Description",
		Status:      "pending",
		Priority:    "medium",
		DueDate:     time.Now(),
	}
	err := db.Create(&task).Error
	assert.Nil(t, err)

	url := fmt.Sprintf("/tasks/%d", task.ID)

	updateRequest := map[string]interface{}{
		"title": "",
	}

	apitest.
		Handler(app.Router).
		Put(url).
		Header("Authorization", "Bearer "+token).
		JSON(updateRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		Body(`{"ok": false, "error": "title is required"}`).
		End()
}
