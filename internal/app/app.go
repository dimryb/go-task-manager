package app

import (
	"github.com/gorilla/mux"
	"go-task-manager/internal/infrastructure/config"
	"gorm.io/gorm"
	"net/http"
)

type App struct {
	Config config.Config
	Router *mux.Router
	DB     *gorm.DB
}

func NewApp() App {
	return App{}
}

func (app *App) Setup() error {
	router := mux.NewRouter()

	app.Router = router

	return nil
}

func (app *App) Run() error {
	return http.ListenAndServe(":8080", app.Router)
}
