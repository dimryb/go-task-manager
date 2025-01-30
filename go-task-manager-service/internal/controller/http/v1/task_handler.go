package v1

import "net/http"

type TaskHandler interface {
	CreateTask(writer http.ResponseWriter, request *http.Request)
	GetTaskById(writer http.ResponseWriter, request *http.Request) //todo: не требуется
	GetTasksFiltered(writer http.ResponseWriter, request *http.Request)
	UpdateTask(writer http.ResponseWriter, request *http.Request)
	DeleteTask(writer http.ResponseWriter, request *http.Request)

	ExportTasks(w http.ResponseWriter, r *http.Request)
	ImportTasks(w http.ResponseWriter, r *http.Request)
}
