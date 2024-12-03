package app

import (
	"context"
	service "github.com/kavehrafie/go-scheduler/internal/service/scheduler"
	"github.com/kavehrafie/go-scheduler/internal/store/sqlite"
	"github.com/kavehrafie/go-scheduler/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type App struct {
	server       *echo.Echo
	SchedulerSvc *service.Scheduler
}

func NewApp(log *logrus.Logger) *App {

	dbCfg := database.Config{
		Driver: database.SQLite,
		URL:    "./db/sql.db",
	}

	factory := &sqlite.SQLiteFactory{}
	store, err := factory.NewStore(dbCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	//svc := NewSchedulerService(store)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	scheduler := service.NewScheduler(&store, log)
	app := &App{
		server:       e,
		SchedulerSvc: scheduler,
	}

	app.routes()

	return app
}

func (a *App) routes() {
	//api := a.server.Group("/api")

	//api.POST("/scheduled-actions", a.svc)
}

func (a *App) Start(addr string) error {
	return a.server.Start(addr)
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
