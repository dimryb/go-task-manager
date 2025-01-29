package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-task-manager/internal/controller/http/rest"
	"net/http"
	"strings"
)

var jwtSecret = []byte("secret_key")

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			rest.WriteError(w, http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			rest.WriteError(w, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
