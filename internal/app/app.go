package app

import (
	"github.com/kavehrafie/go-scheduler/internal/store/sqlite"
	"github.com/kavehrafie/go-scheduler/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type App struct {
	server *echo.Echo
	//svc    *SchedulerService
}

func NewApp() *App {

	dbCfg := database.Config{
		Driver:   database.SQLite,
		Database: "./db/basket.db",
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

	app := &App{
		server: e,
		//svc: svc
	}

	app.routes()

	return app
}

func (a *App) routes() {
	api := a.server.Group("/api")

	//api.POST("/scheduled-actions", a.svc)
}
