package api

import (
	"net/http"

	"github.com/MauricioGZ/CRUD-GO/internal/api/dtos"
	"github.com/MauricioGZ/CRUD-GO/internal/service"
	"github.com/labstack/echo/v4"
)

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	uParams := dtos.RegisterUser{}

	if err := c.Bind(&uParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	if err := a.serv.RegisterUser(
		ctx,
		uParams.FirstName,
		uParams.LastName,
		uParams.Email,
		uParams.Password,
	); err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusBadRequest, responseMessage{Message: "user already exists"})
		}
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}
	return c.JSON(http.StatusCreated, responseMessage{Message: "user created successfully"})
}

func (a *API) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	uParams := dtos.LoginUser{}

	if err := c.Bind(&uParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	u, err := a.serv.LoginUser(ctx, uParams.Email, uParams.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, u)

}
