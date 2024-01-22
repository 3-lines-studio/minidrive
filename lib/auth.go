package lib

import "github.com/labstack/echo/v4"

type AppContext struct {
	echo.Context
	UserId string
}

func GetUserId(c echo.Context) string {
	appContext := c.(*AppContext)

	return appContext.UserId
}
