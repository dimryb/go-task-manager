package router

import (
	"github.com/gorilla/mux"
	"net/http"

	"go-task-manager/internal/infrastructure/http/handler"
)

func NewRouter(taskHandler handler.TaskHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(`/tasks/{id}`, taskHandler.GetTaskById).Methods(http.MethodGet)
	r.HandleFunc(`/tasks`, taskHandler.CreateTask).Methods(http.MethodPost)
	return r
}
