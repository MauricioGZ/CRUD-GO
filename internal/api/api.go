package api

import (
	"context"
	"fmt"

	"github.com/MauricioGZ/CRUD-GO/internal/service"
	"github.com/labstack/echo/v4"
)

type API struct {
	serv service.Service
}

type responseMessage struct {
	Message string `json:"message"`
}

func New(_serv service.Service) *API {
	return &API{
		serv: _serv,
	}
}

func (a *API) Start(e *echo.Echo, address string) error {
	ctx := context.Background()

	err := a.serv.GetAllPermissionsRoles(ctx)
	if err != nil {
		fmt.Println(err)
	}

	err = a.serv.GetAllRoles(ctx)
	if err != nil {
		fmt.Println(err)
	}

	a.RegisterRoutes(e)
	return e.Start((fmt.Sprintf(":%s", address)))
}
