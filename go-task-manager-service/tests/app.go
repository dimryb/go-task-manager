package tests

import (
	"testing"

	"go-task-manager-service/internal/app"
	"go-task-manager-service/internal/repo/pgdb"
)

func AppSetup(t *testing.T) *app.App {
	app := app.Initialize(".env.test")

	pgdb.MigrateUp(app.DB)
	return app
}

func AppTeardown(app *app.App) {
	pgdb.MigrateDown(app.DB)
	app.Teardown()
}
