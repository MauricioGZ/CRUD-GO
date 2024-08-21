package main

import (
	"log"

	"github.com/MauricioGZ/CRUD-GO/internal/api"
	"github.com/MauricioGZ/CRUD-GO/internal/db"
	"github.com/MauricioGZ/CRUD-GO/internal/repository"
	"github.com/MauricioGZ/CRUD-GO/internal/service"
	"github.com/MauricioGZ/CRUD-GO/settings"
	"github.com/labstack/echo/v4"
)

func main() {
	s, err := settings.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.New(*s)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.New(db)
	serv := service.New(repo)
	a := api.New(serv)
	e := echo.New()
	a.Start(e, s.Port)
}
