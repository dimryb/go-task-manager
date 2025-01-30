package pgdb

import "go-task-manager-service/internal/repo/pgdb/models"

type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
}
