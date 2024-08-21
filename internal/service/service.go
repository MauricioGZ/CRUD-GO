package service

import (
	"context"

	"github.com/MauricioGZ/CRUD-GO/internal/model"
	"github.com/MauricioGZ/CRUD-GO/internal/repository"
)

type Service interface {
	RegisterUser(ctx context.Context, firstName, lastName, email, password string) error
	LoginUser(ctx context.Context, email, password string) (*model.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(_repo repository.Repository) Service {
	return &serv{
		repo: _repo,
	}
}
