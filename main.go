package main

import (
	"github.com/3-lines-studio/minidrive/handler"
	"github.com/3-lines-studio/minidrive/lib"
	"github.com/carbonyde/tungsten"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
)

func main() {
	goth.UseProviders(
		github.New(
			lib.Env.GithubClientID,
			lib.Env.GithubClientSecret,
			"http://localhost:3000/auth/github/callback",
		),
	)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")

	e.Use(handler.AuthMiddleware)

	handler.Auth(e)
	handler.View(e)

	if lib.Env.Watch {
		e.GET("/hot-reload", tungsten.HotReload)
	}

	e.Logger.Fatal(e.Start("0.0.0.0:3000"))
}
