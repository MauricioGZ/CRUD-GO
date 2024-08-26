package api

import (
	"github.com/labstack/echo/v4"
)

func (a *API) RegisterRoutes(e *echo.Echo) {
	user := e.Group("/user")
	address := user.Group("/address")

	user.POST("/register", a.RegisterUser)
	user.GET("/login", a.LoginUser)

	address.GET("/", a.GetAddresses)
	address.POST("/register", a.AddAddress)
	address.POST("/update", a.UpdateAddress)
	address.POST("/delete/:id", a.DeleteAddress)
}
