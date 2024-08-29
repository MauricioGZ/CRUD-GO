package api

import (
	"net/http"

	"github.com/MauricioGZ/CRUD-GO/internal/api/dtos"
	"github.com/MauricioGZ/CRUD-GO/internal/service"
	"github.com/labstack/echo/v4"
)

func (a *API) AddProduct(c echo.Context) error {
	ctx := c.Request().Context()
	pParams := dtos.RegisterProduct{}
	_, err := validateUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "unauthorized"})
	}

	//TODO: add permissions validation

	if err := c.Bind(&pParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	err = a.serv.RegisterProduct(
		ctx,
		pParams.Name,
		pParams.Description,
		pParams.Price,
		pParams.Stock,
		pParams.CategoryID,
		pParams.Image,
	)

	if err != nil {
		if err == service.ErrCategoryDoesNotExist {
			return c.JSON(http.StatusBadRequest, responseMessage{Message: "category does not exist"})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "product added successfully"})
}
