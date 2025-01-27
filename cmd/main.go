package main

import (
	"go-task-manager/internal/app"
	"go-task-manager/internal/infrastructure/config"
	"log"
)

func main() {
	cfg := config.Config{}
	if err := cfg.Load(".env"); err != nil {
		panic(err)
	}

	application := app.NewApp(&cfg)
	application.Initialize()

	log.Println("Application is starting...")
	if err := application.Run(); err != nil {
		log.Fatalf("Application failed to start: %v", err)
	}
}
