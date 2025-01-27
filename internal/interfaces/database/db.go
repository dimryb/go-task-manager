package database

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect to database: " + err.Error())
	}

	// todo: добавить миграцию

	log.Println("Database connection established successfully")
	return db, nil
}
