package service

import (
	"errors"
	repository "go-task-manager-service/internal/repo/pgdb"
	"go-task-manager-service/internal/repo/pgdb/models"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret_key")

type authUseCase struct {
	UserRepo repository.UserRepository
}

func NewAuthUseCase(userRepo repository.UserRepository) AuthUseCase {
	return &authUseCase{
		UserRepo: userRepo,
	}
}

func (u *authUseCase) Register(username, password string) error {
	user := models.User{Username: username, Password: password}
	return u.UserRepo.CreateUser(&user)
}

func (u *authUseCase) Login(username, password string) (string, error) {
	user, err := u.UserRepo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
