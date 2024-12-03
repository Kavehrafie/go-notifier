package main

import (
	"context"
	"github.com/kavehrafie/go-scheduler/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})

	// ✅ load config
	//_, err := config.Load()
	//if err != nil {
	//	logger.Fatal(err)
	//}

	a := app.NewApp(logger)

	// create context fpr scheduler
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start scheduler
	if err := a.SchedulerSvc.Start(ctx); err != nil {
		logger.Fatal(err)
	}
	defer a.SchedulerSvc.Stop()

	//go func() {
	if err := a.Start(":8080"); err != nil {
		logger.Fatal(err)
	}
	//}()

	// graceful shutdown
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//
	//ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//
	//if err := a.Shutdown(ctx); err != nil {
	//	logger.Error(err)
	//}
}
