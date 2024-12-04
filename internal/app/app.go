package app

import (
	"context"
	"fmt"
	"github.com/kavehrafie/go-scheduler/internal/app/handlers"
	"github.com/kavehrafie/go-scheduler/internal/repository"
	"github.com/kavehrafie/go-scheduler/internal/service"
	"github.com/kavehrafie/go-scheduler/pkg/config"
	"github.com/kavehrafie/go-scheduler/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Log   *logrus.Logger
	Cfg   *config.Config
	Echo  *echo.Echo
	Repo  *repository.Repository
	ScSvc *service.SchedulerService
}

func New() *App {
	// init logger
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	// load config
	cfg := loadConfig()

	// init database
	db, err := database.NewStore(cfg, logrus.New())
	if err != nil {
		log.Fatal(err)
	}

	repo, err := repository.NewRepository(&db, log)
	if err != nil {
		log.Fatal(err)
	}

	svc := service.NewSchedulerService(repo.GetTaskRepository(), log)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := handlers.NewHandler(repo, log)
	h.RegisterRoutes(e)

	return &App{
		Log:   log,
		Cfg:   cfg,
		Echo:  e,
		Repo:  &repo,
		ScSvc: svc,
	}
}

func loadConfig() *config.Config {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config")
	}

	return cfg
}

func (a *App) Run() error {
	a.ScSvc.Start()
	defer a.ScSvc.Stop()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := a.Echo.Start(fmt.Sprintf(":%v", a.Cfg.Server.Port)); err != nil && err != http.ErrServerClosed {
			a.Log.Fatalf("shutting down the server: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-quit

	// Gracefully shutdown the server with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.Echo.Shutdown(ctx); err != nil {
		return fmt.Errorf("error during server shutdown: %v", err)
	}

	return nil
}
