package models

import "time"

type Task struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:255"`
	Status    string `gorm:"size:50"`
	Priority  string `gorm:"size:50"`
	DueDate   time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
