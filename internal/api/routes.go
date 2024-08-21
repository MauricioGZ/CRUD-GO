package api

import (
	"github.com/labstack/echo/v4"
)

func (a *API) RegisterRoutes(e *echo.Echo) {
	user := e.Group("/user")

	user.POST("/register", a.RegisterUser)
	user.GET("/login", a.LoginUser)
}
