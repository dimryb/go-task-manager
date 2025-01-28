package app

import (
	"github.com/gorilla/mux"
	"go-task-manager/config"
	"go-task-manager/internal/controller/http/v1"
	"go-task-manager/internal/repo/pgdb"
	"gorm.io/gorm"
	"log"
	"net/http"

	"go-task-manager/internal/service"
)

type App struct {
	Config *config.Config
	Router *mux.Router
	DB     *gorm.DB
}

func NewApp(cfg *config.Config) *App {
	return &App{
		Config: cfg,
	}
}

func (app *App) Initialize() {
	app.setupDatabase()
	app.setupRouter()
}

func (app *App) setupDatabase() {
	db, err := pgdb.Connect(app.Config.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	app.DB = db
}

func (app *App) setupRouter() {
	taskRepo := pgdb.NewTaskRepository(app.DB)
	taskUseCase := service.NewTaskUseCase(taskRepo)
	taskHandler := v1.NewTaskHandler(taskUseCase)
	app.Router = v1.NewRouter(taskHandler)
}

func (app *App) Run() error {
	return http.ListenAndServe(":8080", app.Router)
}

func (app *App) Teardown() error {
	db, err := app.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
