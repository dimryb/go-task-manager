package v1

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(taskHandler TaskHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(`/tasks/{id}`, taskHandler.GetTaskById).Methods(http.MethodGet)
	r.HandleFunc(`/tasks`, taskHandler.CreateTask).Methods(http.MethodPost)
	return r
}
