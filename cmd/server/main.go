package main

import (
	"context"
	"github.com/kavehrafie/go-scheduler/internal/config"
	"github.com/kavehrafie/go-scheduler/internal/store/sqlite"
	"github.com/kavehrafie/go-scheduler/pkg/database"
	"log"
	"time"
)

func main() {
	// ✅ load config
	_, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// 1. load config

	// 2. load store
	// 2.1 set the repository store
	// 2.2 set up the notification service

	cfg := database.Config{
		Driver: database.SQLite,
		URL:    "./schedules.db",
	}

	factory := &sqlite.SQLiteFactory{}
	store, err := factory.NewStore(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	//sa := &model.ScheduledAction{
	//	ID:          uuid.New().String(),
	//	Title:       "Example Action",
	//	Status:      model.StatusPending,
	//	URL:         "http://example.com",
	//	ScheduledAt: time.Now().Add(25 * time.Hour),
	//}

	ctx := context.Background()
	//if err := store.Create(ctx, sa); err != nil {
	//	log.Fatal(err)
	//}

	pending, err := store.ListPending(ctx, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(pending)

	// ✅ 3. echo server
	//e := echo.New()
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// 4. setup routes

	// 5. logger
	//e.Logger.Fatal(e.Start(":8080"))

}
