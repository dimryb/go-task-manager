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

	"go-task-manager-service/config"
	handler "go-task-manager-service/internal/controller/http/v1"
	repository "go-task-manager-service/internal/repo/pgdb"
	"go-task-manager-service/internal/service"

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
	db, err := repository.Connect(app.Config.PG.URL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	app.DB = db
}

func (app *App) setupRouter() {
	userRepo := repository.NewUserRepository(app.DB)
	authUseCase := service.NewAuthUseCase(userRepo)
	authHandler := handler.NewAuthHandler(authUseCase)

	taskRepo := repository.NewTaskRepository(app.DB)
	taskUseCase := service.NewTaskUseCase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUseCase)
	app.Router = handler.NewRouter(authHandler, taskHandler)
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

func (app *App) initTaskCleanupScheduler() {
	stop := make(chan struct{})

	service.StartTaskCleanupScheduler(stop)
}

func Run(configPath string) {
	app := Initialize(".env")

	app.initTaskCleanupScheduler()

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
