package v1

import (
	"github.com/gorilla/mux"
	"net/http"

	swagger "github.com/swaggo/http-swagger"
	_ "go-task-manager/docs"
)

func NewRouter(taskHandler TaskHandler) *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(swagger.WrapHandler)

	r.HandleFunc(`/tasks/{id}`, taskHandler.GetTaskById).Methods(http.MethodGet)
	r.HandleFunc(`/tasks`, taskHandler.CreateTask).Methods(http.MethodPost)
	r.HandleFunc(`/tasks`, taskHandler.GetTasksFiltered).Methods(http.MethodGet)
	return r
}
