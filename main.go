package main

import (
	"github.com/3-lines-studio/minidrive/handler"
	"github.com/3-lines-studio/minidrive/lib"
	"github.com/carbonyde/tungsten"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")

	handler.View(e)

	if lib.Env.Watch {
		e.GET("/hot-reload", tungsten.HotReload)
	}

	e.Logger.Fatal(e.Start("0.0.0.0:3000"))
}
