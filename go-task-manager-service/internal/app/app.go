package app

import (
	"github.com/gorilla/mux"
	"go-task-manager/config"
	"go-task-manager/internal/controller/http/v1"
	"go-task-manager/internal/repo/pgdb"
	"go-task-manager/internal/service"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type App struct {
	Config *config.Config
	Router *mux.Router
	DB     *gorm.DB
}

func newApp(cfg *config.Config) *App {
	return &App{
		Config: cfg,
	}
}

func (app *App) setupDatabase() {
	db, err := pgdb.Connect(app.Config.PG.URL)
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

func (app *App) Teardown() error {
	db, err := app.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func Initialize(configName string) *App {
	// Configuration
	cfg, err := config.NewConfig(configName)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app := newApp(cfg)

	app.setupDatabase()
	app.setupRouter()

	return app
}

func Run() {
	app := Initialize(".env")

	log.Println("Application is starting...")
	if err := http.ListenAndServe(":8080", app.Router); err != nil {
		log.Fatalf("Application failed to start: %v", err)
	}
}
