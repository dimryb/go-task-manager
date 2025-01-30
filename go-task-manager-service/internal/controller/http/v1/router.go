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

	protected := r.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/tasks/export", taskHandler.ExportTasks).Methods(http.MethodGet)
	protected.HandleFunc("/tasks/import", taskHandler.ImportTasks).Methods(http.MethodPost)

	protected.HandleFunc(`/tasks/{id}`, taskHandler.GetTaskById).Methods(http.MethodGet)
	protected.HandleFunc(`/tasks`, taskHandler.CreateTask).Methods(http.MethodPost)
	protected.HandleFunc(`/tasks`, taskHandler.GetTasksFiltered).Methods(http.MethodGet)
	protected.HandleFunc(`/tasks/{id}`, taskHandler.UpdateTask).Methods(http.MethodPut)
	protected.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods(http.MethodDelete)

	return r
}
