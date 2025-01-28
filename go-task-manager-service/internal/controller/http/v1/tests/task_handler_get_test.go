package tests

import (
	"fmt"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/assert"
	"go-task-manager/internal/repo/pgdb/models"
	"net/http"
	"testing"
	"time"

	"go-task-manager/tests"
)

func TestGetTask_Success(t *testing.T) {
	app := tests.AppSetup(t)
	defer tests.AppTeardown(app)

	db := app.DB

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
		Expect(t).
		Status(http.StatusOK).
		Header("Content-Type", "application/json").
		Assert(
			jsonpath.Chain().
				Equal("ok", true).
				Equal("result.Title", "Test Task").
				Equal("result.Description", "Task for testing").
				Equal("result.Status", "pending").
				Equal("result.Priority", "medium").
				End(),
		).
		End()
}
