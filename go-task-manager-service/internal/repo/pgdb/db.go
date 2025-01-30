package pgdb

import (
	"errors"
	"fmt"
	"go-task-manager-service/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect to database: " + err.Error())
	}

	if err := MigrateUp(db); err != nil {
		return nil, errors.New("failed to run SQL migration: " + err.Error())
	}

	log.Println("Database connection established successfully")
	return db, nil
}

func MigrateUp(db *gorm.DB) error {
	pattern := filepath.Join(utils.GetProjectRoot(), "migrations", "*.up.sql")
	sql, err := ConcatMigrations(pattern)
	if err != nil {
		return err
	}

	err = db.Exec(sql).Error
	if err != nil {
		MigrateDown(db)
		err = db.Exec(sql).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func MigrateDown(db *gorm.DB) error {
	pattern := filepath.Join(utils.GetProjectRoot(), "migrations", "*.down.sql")
	sql, err := ConcatMigrations(pattern)
	if err != nil {
		return err
	}

	err = db.Exec(sql).Error
	if err != nil {
		return err
	}

	return nil
}

func ConcatMigrations(pattern string) (string, error) {
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		return "", err
	}

	if len(filenames) == 0 {
		fmt.Println("Warning: No migration files found.")
		return "", nil
	}

	sort.Strings(filenames)

	var contents []string
	for _, filename := range filenames {
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			return "", err
		}
		contents = append(contents, string(bytes))
	}
	return strings.Join(contents, "\n\n"), nil
}
