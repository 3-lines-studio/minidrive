package handler

import (
	"github.com/3-lines-studio/minidrive/lib"
	"github.com/3-lines-studio/minidrive/view"
	"github.com/labstack/echo/v4"
)

func View(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return lib.Render(c, view.HomePage())
	})
}
