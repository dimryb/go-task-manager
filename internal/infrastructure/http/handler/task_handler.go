package handler

import "net/http"

type TaskHandler interface {
	CreateTask(writer http.ResponseWriter, request *http.Request)
	GetTaskById(writer http.ResponseWriter, request *http.Request)
}
