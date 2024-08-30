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
	_, role, err := validateUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "unauthorized"})
	}

	if err := c.Bind(&pParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	err = a.serv.RegisterProduct(
		ctx,
		role,
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
		if err == service.ErrInvalidPermissions {
			return c.JSON(http.StatusUnauthorized, responseMessage{Message: service.ErrInvalidPermissions.Error()})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "product added successfully"})
}

func (a *API) GetAllProducts(c echo.Context) error {
	ctx := c.Request().Context()

	pParams := dtos.GetProductByCategory{}

	if err := c.Bind(&pParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	if pParams.CategoryName != "" {
		products, err := a.serv.GetProductsByCategory(ctx, pParams.CategoryName)

		if err != nil {
			if err == service.ErrCategoryDoesNotExist {
				return c.JSON(http.StatusBadRequest, responseMessage{Message: service.ErrCategoryDoesNotExist.Error()})
			}
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
		}

		return c.JSON(http.StatusOK, products)
	}

	products, err := a.serv.GetAllProducts(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, products)
}

func (a *API) GetProductByID(c echo.Context) error {
	ctx := c.Request().Context()
	pParams := dtos.GetProductByID{}

	if err := c.Bind(&pParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	product, err := a.serv.GetProductByID(ctx, pParams.ID)

	if err != nil {
		if err == service.ErrProductDoesNotExist {
			return c.JSON(http.StatusBadRequest, responseMessage{Message: service.ErrProductDoesNotExist.Error()})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, product)
}
