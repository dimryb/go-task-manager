package tests

import (
	"github.com/stretchr/testify/require"
	"go-task-manager/config"
	"go-task-manager/internal/repo/pgdb"
	"path/filepath"
	"testing"

	"go-task-manager/internal/app"
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

	pgdb.MigrateUp(app.DB)
	return app
}

func AppTeardown(app *app.App) {
	pgdb.MigrateDown(app.DB)
	app.Teardown()
}
