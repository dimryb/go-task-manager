package app

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"

	"go-task-manager/internal/infrastructure/config"
	"go-task-manager/internal/infrastructure/database"
	"go-task-manager/internal/infrastructure/http/handler"
	"go-task-manager/internal/infrastructure/repository"
	"go-task-manager/internal/infrastructure/router"
	"go-task-manager/internal/usecase"
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
	db, err := database.Connect(app.Config.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	app.DB = db
}

func (app *App) setupRouter() {
	taskRepo := repository.NewTaskRepository(app.DB)
	taskUseCase := usecase.NewTaskUseCase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUseCase)
	app.Router = router.NewRouter(taskHandler)
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
