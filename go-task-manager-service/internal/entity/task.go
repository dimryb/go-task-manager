package entity

import "time"

type Task struct {
	ID          uint
	Title       string
	Description string
	Status      string
	Priority    string
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
