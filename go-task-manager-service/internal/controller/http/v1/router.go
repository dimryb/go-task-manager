package v1

import (
	"github.com/gorilla/mux"
	"go-task-manager/internal/controller/http/v1/middleware"
	"net/http"

	swagger "github.com/swaggo/http-swagger"
	_ "go-task-manager/docs"
)

func NewRouter(authHandler AuthHandler, taskHandler TaskHandler) *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(swagger.WrapHandler)

	r.HandleFunc("/auth/register", authHandler.Register).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", authHandler.Login).Methods(http.MethodPost)

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTMiddleware)

	r.HandleFunc(`/tasks/{id}`, taskHandler.GetTaskById).Methods(http.MethodGet)
	r.HandleFunc(`/tasks`, taskHandler.CreateTask).Methods(http.MethodPost)
	r.HandleFunc(`/tasks`, taskHandler.GetTasksFiltered).Methods(http.MethodGet)
	r.HandleFunc(`/tasks/{id}`, taskHandler.UpdateTask).Methods(http.MethodPut)
	r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods(http.MethodDelete)

	return r
}
