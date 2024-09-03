package api

import (
	"fmt"
	"net/http"

	"github.com/MauricioGZ/CRUD-GO/internal/api/dtos"
	"github.com/MauricioGZ/CRUD-GO/internal/service"
	"github.com/labstack/echo/v4"
)

func (a *API) AddOrder(c echo.Context) error {
	oParams := []dtos.OrderItem{}
	var totalPrice float32 = 0.0
	ctx := c.Request().Context()
	email, _, err := validateUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "unauthorized"})
	}

	if err := c.Bind(&oParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	for _, o := range oParams {
		totalPrice += (o.Price * float32(o.Quantity))
	}

	order, err := a.serv.RegisterOrder(ctx, email, "pending", totalPrice)

	if err != nil {
		if err == service.ErrUserDoesntExist {
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: service.ErrUserDoesntExist.Error()})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	for _, o := range oParams {
		err = a.serv.RegisterOrderItem(ctx, order.ID, o.ProductID, o.Quantity, o.Price)
		if err != nil {
			if err == service.ErrProductDoesNotExist {
				return c.JSON(http.StatusBadRequest, responseMessage{Message: fmt.Sprintf("product_id:%d does not exist", o.ProductID)})
			}
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
		}
	}

	return c.JSON(http.StatusOK, order)
}

func (a *API) GetOrders(c echo.Context) error {
	ctx := c.Request().Context()
	email, _, err := validateUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "unauthorized"})
	}

	orders, err := a.serv.GetOrdersByUser(ctx, email)

	if err != nil {
		if err == service.ErrUserDoesntExist {
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: service.ErrUserDoesntExist.Error()})
		}
		if err == service.ErrNoOrdersRegistered {
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: service.ErrNoOrdersRegistered.Error()})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, orders)
}
