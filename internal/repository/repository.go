package repository

import (
	"context"
	"database/sql"

	"github.com/MauricioGZ/CRUD-GO/internal/entity"
)

// The Repository interfaces wraps the CRUD operations
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	//user interfaces
	SaveUser(ctx context.Context, firstName, lastName, email, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	DeleteUserByEmail(ctx context.Context, email string) error
	//address interfaces
	SaveAddress(ctx context.Context, userId int64, addressType, address, city, state, country, zipCode string) error
	GetAddressesByUserId(ctx context.Context, userId int64) ([]entity.Address, error)
	DeleteAddressByID(ctx context.Context, id int64) error
	UpdateAddressByID(ctx context.Context, id int64, addressType, address, city, state, country, zipCode string) error
}

type repo struct {
	db *sql.DB
}

func New(_db *sql.DB) Repository {
	return &repo{
		db: _db,
	}
}
