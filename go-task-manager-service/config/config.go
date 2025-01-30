package config

import (
	"errors"
	"github.com/joho/godotenv"
	"go-task-manager-service/pkg/utils"
	"os"
	"path/filepath"
)

type (
	Config struct {
		HTTP
		Log
		PG
	}

	HTTP struct {
		Port string
	}

	PG struct {
		URL string
	}

	Log struct {
		Level string
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

	cfg.Log.Level = os.Getenv("LOG_LEVEL")
	if cfg.Log.Level == "" {
		return nil, errors.New("missing LOG_LEVEL")
	}

	cfg.HTTP.Port = os.Getenv("HTTP_PORT")
	if cfg.HTTP.Port == "" {
		return nil, errors.New("missing HTTP_PORT")
	}

	return &cfg, nil
}
