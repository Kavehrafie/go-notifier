package main

import (
	"context"
	"github.com/kavehrafie/go-scheduler/internal/app"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := app.New()

	// handle shutdown gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		logrus.Infof("received signal: %v", sig)
		cancel()
	}()

	if err := application.Run(ctx); err != nil {
		logrus.Fatalf("error running application: %v", err)
	}
}
