package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-task-manager/config"
	"go-task-manager/internal/controller/http/v1"
	"go-task-manager/internal/repo/pgdb"
	"go-task-manager/internal/service"

	log "github.com/sirupsen/logrus"
)

type App struct {
	Config *config.Config
	Router *mux.Router
	DB     *gorm.DB
	Logger *log.Logger
}

func newApp(cfg *config.Config) *App {
	return &App{
		Config: cfg,
		Logger: log.New(),
	}
}

func (app *App) setupLogger() {
	level, err := log.ParseLevel(app.Config.Log.Level)
	if err != nil {
		app.Logger.Fatalf("Failed to parse log level: %v", err)
	}
	app.Logger.SetLevel(level)
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
	app.setupLogger()
	app.setupDatabase()
	app.setupRouter()

	return app
}

func Run(configPath string) {
	app := Initialize(".env")

	// HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.HTTP.Port),
		Handler: app.Router,
	}

	// Start server
	go func() {
		app.Logger.Infof("Application is starting on port %s...", app.Config.HTTP.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatalf("Application failed to start: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	app.Logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		app.Logger.Fatalf("Server forced to shutdown: %v", err)
	}

	app.Logger.Info("Server exiting")
}
