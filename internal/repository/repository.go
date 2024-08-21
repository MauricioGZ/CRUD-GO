package repository

import (
	"context"
	"database/sql"

	"github.com/MauricioGZ/CRUD-GO/internal/entity"
)

type Repository interface {
	SaveUser(ctx context.Context, firstName, lastName, email, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type repo struct {
	db *sql.DB
}

func New(_db *sql.DB) Repository {
	return &repo{
		db: _db,
	}
}
