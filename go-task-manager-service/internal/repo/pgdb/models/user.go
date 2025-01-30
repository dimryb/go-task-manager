package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"unique;not null"`
	Password  string `json:"-" gorm:"not null"` // Хранить зашифрованным
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
