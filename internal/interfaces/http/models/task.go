package models

type CreateTaskRequest struct {
	Title    string `json:"title" binding:"required"`
	Status   string `json:"status" binding:"required,oneof=pending in_progress done"`
	Priority string `json:"priority" binding:"required,oneof=low medium high"`
	DueDate  string `json:"due_date" binding:"required"`
}

type TaskResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
	DueDate  string `json:"due_date"`
}
