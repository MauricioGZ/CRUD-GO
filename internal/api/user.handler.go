package api

import (
	"net/http"

	"github.com/MauricioGZ/CRUD-GO/encryption"
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
		"Customer",
	); err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusBadRequest, responseMessage{Message: "user already exists"})
		}
		if err == service.ErrRolesNotInitialized {
			return c.JSON(http.StatusInternalServerError, responseMessage{Message: service.ErrRolesNotInitialized.Error()})
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

	user, err := a.serv.LoginUser(ctx, uParams.Email, uParams.Password)

	if err != nil {
		if err == service.ErrInvalidCredentials {
			return c.JSON(http.StatusBadRequest, responseMessage{Message: service.ErrInvalidCredentials.Error()})
		}
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "internal server error"})
	}

	token, err := encryption.SignedLoginToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseMessage{Message: "intertnal server error"})
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, user)
}

func validateUser(c echo.Context) (string, error) {
	cookie, err := c.Cookie("Authorization")
	if err != nil {
		return "", err
	}

	claims, err := encryption.ParseLoginJWT(cookie.Value)
	if err != nil {
		return "", err
	}

	return claims["email"].(string), nil
}
