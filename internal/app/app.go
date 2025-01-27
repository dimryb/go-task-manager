package app

import (
	"github.com/gorilla/mux"
	"go-task-manager/internal/interfaces/database"
	"go-task-manager/internal/interfaces/repository"
	"go-task-manager/internal/usecase"
	"gorm.io/gorm"
	"log"
	"net/http"

	"go-task-manager/internal/infrastructure/config"
	"go-task-manager/internal/infrastructure/router"
	"go-task-manager/internal/interfaces/http/handler"
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
