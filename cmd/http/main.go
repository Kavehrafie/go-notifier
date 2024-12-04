package main

import (
	"fmt"
	"github.com/kavehrafie/go-scheduler/internal/app"
)

func main() {
	a := app.New()

	if err := a.Run(); err != nil {
		fmt.Println(err)
	}
	//port := fmt.Sprintf(":%s", a.Cfg.Server.Port)

	// Graceful shutdown
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, os.Interrupt)
	//go func() {
	//	<-quit
	//	if err := a.Echo.Shutdown(context.Background()); err != nil {
	//		a.Echo.Logger.Fatal(err)
	//	}
	//}()

	//a.Echo.Logger.Fatal(a.Echo.Start(port))
}
