package api

import (
	"net/http"

	"github.com/MauricioGZ/CRUD-GO/internal/service"
	"github.com/labstack/echo/v4"
)

func (a *API) GetOrderItemsByUser(c echo.Context) error {
	email, _, err := validateUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "unauthorized"})
	}

	ctx := c.Request().Context()

	orderItems, err := a.serv.GetOrderItemsByUser(ctx, email)

	if err != nil {
		if err == service.ErrUserDoesntExist {
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: service.ErrUserDoesntExist.Error()})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, orderItems)
}
