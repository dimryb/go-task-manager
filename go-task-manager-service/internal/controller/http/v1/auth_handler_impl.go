package v1

import (
	"encoding/json"
	"net/http"

	"go-task-manager-service/internal/controller/http/models"
	"go-task-manager-service/internal/controller/http/rest"
	"go-task-manager-service/internal/service"
)

type authHandler struct {
	AuthUseCase service.AuthUseCase
}

func NewAuthHandler(authUseCase service.AuthUseCase) AuthHandler {
	return &authHandler{AuthUseCase: authUseCase}
}

// Register godoc
// @Summary Регистрация пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.RegisterUserRequest true "Данные пользователя"
// @Success 201 {object} rest.Response
// @Failure 400 {object} rest.Response
// @Router /auth/register [post]
func (h authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rest.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.AuthUseCase.Register(req.Username, req.Password)
	if err != nil {
		rest.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	rest.WriteJSON(w, http.StatusCreated, rest.Response{Ok: true})
}

// Login godoc
// @Summary Авторизация пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.LoginUserRequest true "Данные пользователя"
// @Success 200 {object} rest.Response
// @Failure 401 {object} rest.Response
// @Router /auth/login [post]
func (h authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rest.WriteError(w, http.StatusBadRequest, err)
		return
	}

	token, err := h.AuthUseCase.Login(req.Username, req.Password)
	if err != nil {
		rest.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	rest.WriteJSON(w, http.StatusOK, rest.Response{Ok: true, Result: token})
}
