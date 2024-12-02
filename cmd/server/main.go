package main

import (
	"github.com/kavehrafie/go-scheduler/internal/config"
	"log"
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

	// ✅ 3. echo server
	//e := echo.New()
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// 4. setup routes

	// 5. logger
	//e.Logger.Fatal(e.Start(":8080"))

}
