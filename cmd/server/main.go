package main

import (
	"github.com/kavehrafie/go-notify/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	// ✅ load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// 1. load config

	// 2. load store
	// 2.1 set the repository store
	// 2.2 setup the notification service

	// ✅ 3. echo server
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// 4. setup routes

	// 5. logger
}
