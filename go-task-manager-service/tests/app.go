package tests

import (
	"go-task-manager/internal/repo/pgdb"
	"testing"

	"go-task-manager/internal/app"
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
