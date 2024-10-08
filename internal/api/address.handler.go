package api

import (
	"net/http"

	"github.com/MauricioGZ/CRUD-GO/internal/api/dtos"
	"github.com/labstack/echo/v4"
)

func (a *API) AddAddress(c echo.Context) error {
	ctx := c.Request().Context()
	aParams := dtos.RegisterAddress{}
	email, _, err := validateUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "unauthorized"})
	}

	if err := c.Bind(&aParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	err = a.serv.RegisterAddress(
		ctx,
		email,
		aParams.AddressType,
		aParams.Address,
		aParams.City,
		aParams.State,
		aParams.Country,
		aParams.ZipCode,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "address registered"})
}

func (a *API) UpdateAddress(c echo.Context) error {
	ctx := c.Request().Context()
	aParams := dtos.UpdateAddress{}
	email, _, err := validateUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "unauthorized"})
	}

	if err := c.Bind(&aParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	err = a.serv.UpdateAddress(
		ctx,
		aParams.ID,
		email,
		aParams.AddressType,
		aParams.Address,
		aParams.City,
		aParams.State,
		aParams.Country,
		aParams.ZipCode,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "address updated"})
}

func (a *API) GetAddresses(c echo.Context) error {
	ctx := c.Request().Context()
	email, _, err := validateUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "unauthorized"})
	}

	aa, err := a.serv.GetAllAddresses(ctx, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, aa)
}

func (a *API) DeleteAddress(c echo.Context) error {
	ctx := c.Request().Context()
	aParams := dtos.DeleteAddress{}
	email, _, err := validateUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "unauthorized"})
	}

	if err := c.Bind(&aParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	err = a.serv.DeleteAddress(ctx, aParams.ID, email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "address successfully deleted"})
}
