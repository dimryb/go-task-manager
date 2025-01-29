package pgdb

import "go-task-manager/internal/repo/pgdb/models"

type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
}
