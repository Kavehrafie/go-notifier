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
	"time"
)

type App struct {
	log       *logrus.Logger
	cfg       *config.Config
	echo      *echo.Echo
	repo      *repository.Repository
	scheduler *service.SchedulerService
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
		log:       log,
		cfg:       cfg,
		echo:      e,
		repo:      &repo,
		scheduler: svc,
	}
}

func loadConfig() *config.Config {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config")
	}

	return cfg
}

func (a *App) Run(ctx context.Context) error {

	a.scheduler.Start(ctx)
	defer a.scheduler.Stop()

	// start the http server
	go func() {
		if err := a.echo.Start(fmt.Sprintf(":%v", a.cfg.Server.Port)); err != nil && err != http.ErrServerClosed {
			a.log.Fatalf("failed to start http server: %v", err)
		}
	}()

	// wait for context cancellation
	<-ctx.Done()

	// graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.echo.Shutdown(shutdownCtx); err != nil {
		a.log.Errorf("failed to shutdown http server: %v", err)
	}

	return nil
}
