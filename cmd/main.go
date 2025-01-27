package main

import (
	"fmt"
	"go-task-manager/internal/app"
)

func main() {
	app := app.NewApp()
	if err := app.Config.Load(".env"); err != nil {
		panic(err)
	}

	if err := app.Setup(); err != nil {
		panic(err)
	}

	fmt.Println("Starting is running.")

	if err := app.Run(); err != nil {
		panic(err)
	}
}
