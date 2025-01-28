package config

import (
	"errors"
	"github.com/joho/godotenv"
	"go-task-manager/pkg/utils"
	"os"
	"path/filepath"
)

type (
	Config struct {
		PG
	}

	PG struct {
		URL string
	}
)

func NewConfig(configName string) (*Config, error) {
	cfg := Config{}
	filename := filepath.Join(utils.GetProjectRoot(), configName)

	err := godotenv.Load(filename)
	if err != nil {
		return nil, errors.New("Error loading " + filename)
	}

	cfg.PG.URL = os.Getenv("DATABASE_URL")
	if cfg.PG.URL == "" {
		return nil, errors.New("missing DATABASE_URL")
	}

	return &cfg, nil
}
