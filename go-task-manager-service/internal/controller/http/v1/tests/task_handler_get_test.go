package tests

import (
	"fmt"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/assert"
	"go-task-manager-service/internal/repo/pgdb/models"
	"net/http"
	"testing"
	"time"

	"go-task-manager-service/tests"
)

func TestGetTask_Success(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	db := app.DB

	token := GetAuthToken(t)
	assert.NotEmpty(t, token)

	task := models.Task{
		Title:       "Test Task",
		Description: "Task for testing",
		Status:      "pending",
		Priority:    "medium",
		DueDate:     time.Now(),
	}
	err := db.Create(&task).Error
	assert.Nil(t, err)

	url := fmt.Sprintf("/tasks/%d", task.ID)

	apitest.
		Handler(app.Router).
		Get(url).
		Header("Authorization", "Bearer "+token).
		Expect(t).
		Status(http.StatusOK).
		Header("Content-Type", "application/json").
		Assert(
			jsonpath.Chain().
				Equal("ok", true).
				Equal("result.title", "Test Task").
				Equal("result.description", "Task for testing").
				Equal("result.status", "pending").
				Equal("result.priority", "medium").
				End(),
		).
		End()
}
