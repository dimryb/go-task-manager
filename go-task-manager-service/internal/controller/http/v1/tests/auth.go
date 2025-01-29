package tests

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var jwtSecret = []byte("secret_key")

func GetAuthToken(t *testing.T) string {
	username := "testuser"
	password := "password123"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.Nil(t, err)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	assert.Nil(t, err, "пароль должен совпадать")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	assert.Nil(t, err, "ошибка при подписи токена")
	assert.NotEmpty(t, tokenString, "токен не должен быть пустым")

	return tokenString
}
