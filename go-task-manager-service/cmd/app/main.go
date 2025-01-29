package main

import (
	"go-task-manager/internal/app"
)

// @title           Task Manager Service
// @version         1.0
// @description     This is a service for managing tasks.

// @contact.name   Rybakov Dmitry
// @contact.email  dimryb@bk.ru

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description					JWT token

func main() {
	app.Run(".env")
}
