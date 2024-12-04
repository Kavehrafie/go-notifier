package main

import (
	"context"
	"database/sql"
	"github.com/kavehrafie/go-scheduler/internal/repository"
	"github.com/kavehrafie/go-scheduler/pkg/config"
	"github.com/kavehrafie/go-scheduler/pkg/database"
	"github.com/kavehrafie/go-scheduler/pkg/domain"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
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

	repos, err := repository.NewRepository(&db, log)
	if err != nil {
		log.Warningln(err)
	}

	ctx := context.Background()
	t := &domain.Task{
		URL:       "https://github.com/kavehrafie/go-scheduler",
		ExecuteAt: sql.NullTime{Time: time.Now().Add(time.Hour), Valid: true},
	}
	tr := repos.GetTaskRepository()
	err = tr.Create(ctx, t, log)
	if err != nil {
		log.Warningln(err)
	}

	log.WithFields(logrus.Fields{
		"taskId": t.ID,
	}).Info("task created")

	tasks, err := tr.ListPendingTasks(ctx)
	if err != nil {
		log.Warningln(err)
	}

	for _, task := range tasks {
		log.WithFields(logrus.Fields{
			"taskId": task.ID,
		}).Info("task pending")
	}

}

func loadConfig() *config.Config {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config")
	}

	return cfg
}
