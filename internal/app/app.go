package app

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type App struct {
	Config Config
	Router *mux.Router
	DB     *gorm.DB
}

type Config struct {
	DatabaseUrl string
}

func (config *Config) Load(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return errors.New("Error loading " + filename)
	}

	config.DatabaseUrl = os.Getenv("DATABASE_URL")
	if config.DatabaseUrl == "" {
		return errors.New("missing DATABASE_URL")
	}

	return nil
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
