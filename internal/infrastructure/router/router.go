package router

import (
	"github.com/gorilla/mux"
	"go-task-manager/internal/interfaces/http/handler"
	"net/http"
)

func NewRouter(taskHandler *handler.TaskHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods(http.MethodPost)
	return r
}
