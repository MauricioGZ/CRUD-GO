package api

import (
	"net/http"

	"github.com/MauricioGZ/CRUD-GO/internal/api/dtos"
	"github.com/MauricioGZ/CRUD-GO/internal/service"
	"github.com/labstack/echo/v4"
)

func (a *API) RegisterCategory(c echo.Context) error {
	ctx := c.Request().Context()
	cParams := dtos.RegisterCategory{}
	_, role, err := validateUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, responseMessage{Message: "unauthorized"})
	}

	if err := c.Bind(&cParams); err != nil {
		return c.JSON(http.StatusBadRequest, responseMessage{Message: "invalid request"})
	}

	if err := a.serv.RegisterCategory(
		ctx,
		role,
		cParams.Name,
		cParams.Description,
		cParams.ParentID,
	); err != nil {
		if err == service.ErrInvalidPermissions {
			return c.JSON(http.StatusBadRequest, responseMessage{Message: service.ErrInvalidPermissions.Error()})
		}
		if err == service.ErrParentCategoryDoesNotExist {
			return c.JSON(http.StatusBadRequest, responseMessage{Message: service.ErrParentCategoryDoesNotExist.Error()})
		}
		if err == service.ErrCategoryAlreadyExists {
			return c.JSON(http.StatusBadRequest, responseMessage{Message: service.ErrCategoryAlreadyExists.Error()})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, responseMessage{Message: "category registered successfully"})
}

func (a *API) GetAllCategories(c echo.Context) error {
	ctx := c.Request().Context()
	categoires, err := a.serv.GetAllCategories(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, categoires)
}
