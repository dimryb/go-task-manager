package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseUrl string
}

func (config *Config) Load(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return errors.New("Error loading " + filename)
	}

	config.DatabaseUrl = os.Getenv("DATABASE_URL")
	if config.DatabaseUrl == "" {
		return errors.New("missing DATABASE_URL")
	}

	return nil
}
