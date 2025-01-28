package tests

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"

	"go-task-manager/internal/app"
	"go-task-manager/internal/infrastructure/config"
	"go-task-manager/internal/infrastructure/database"
	"go-task-manager/pkg/utils"
)

func AppSetup(t *testing.T) *app.App {
	cfg := config.Config{}
	configPath := filepath.Join(utils.GetProjectRoot(), ".env.test")
	err := cfg.Load(configPath)
	if err != nil {
		panic(err)
	}
	require.Nil(t, err)

	app := app.NewApp(&cfg)

	app.Initialize()

	database.MigrateUp(app.DB)
	return app
}

func AppTeardown(app *app.App) {
	database.MigrateDown(app.DB)
	app.Teardown()
}
