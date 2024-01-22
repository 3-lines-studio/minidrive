package handler

import (
	"github.com/3-lines-studio/minidrive/lib"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

var store = sessions.NewCookieStore([]byte(lib.Env.SessionSecret))

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := store.Get(c.Request(), "__session")

		var userId = ""
		sessionUserId, valid := session.Values["user-id"].(string)

		if valid {
			userId = sessionUserId
		}

		cc := &lib.AppContext{Context: c, UserId: userId}

		return next(cc)
	}
}

func Auth(e *echo.Echo) {
	e.GET("/auth/:provider", login)
	e.GET("/auth/:provider/logout", logout)
	e.GET("/auth/:provider/callback", callback)
}

func logout(c echo.Context) error {
	setProvider(c)

	gothic.Logout(c.Response(), c.Request())

	session, _ := store.Get(c.Request(), "__session")
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())

	return c.Redirect(302, "/")
}

func login(c echo.Context) error {
	setProvider(c)

	userId := lib.GetUserId(c)

	if len(userId) > 0 {
		return c.Redirect(302, "/")
	}

	gothic.BeginAuthHandler(c.Response(), c.Request())

	return nil
}

func callback(c echo.Context) error {
	setProvider(c)

	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())

	if err != nil {
		return err
	}

	session, _ := store.Get(c.Request(), "__session")
	// todo: insert user in db and use db id
	session.Values["user-id"] = user.Email
	session.Save(c.Request(), c.Response())

	return c.Redirect(302, "/")
}

func setProvider(c echo.Context) {
	provider := c.Param("provider")
	query := c.Request().URL.Query()
	query.Add("provider", provider)
	c.Request().URL.RawQuery = query.Encode()
}
