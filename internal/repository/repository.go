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
	InsertUser(ctx context.Context, firstName, lastName, email, password string, roleID int64) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	DeleteUserByEmail(ctx context.Context, email string) error
	//address interfaces
	SaveAddress(ctx context.Context, userId int64, addressType, address, city, state, country, zipCode string) error
	GetAddressesByUserId(ctx context.Context, userId int64) ([]entity.Address, error)
	DeleteAddressByID(ctx context.Context, id int64) error
	UpdateAddressByID(ctx context.Context, id int64, addressType, address, city, state, country, zipCode string) error
	//categories interfaces
	InsertCategory(ctx context.Context, name, description string, parentID int64) error
	GetAllCategories(ctx context.Context) ([]entity.Categories, error)
	GetCategoryByID(ctx context.Context, id int64) (*entity.Categories, error)
	GetCategoryByName(ctx context.Context, name string) (*entity.Categories, error)
	UpdateCategoryByID(ctx context.Context, id int64, name, description string, parentID int64) error
	//products interfaces
	InsertProduct(ctx context.Context, name, description string, price float32, stock, categoryId int64, image string) error
	GetAllProducts(ctx context.Context) ([]entity.Product, error)
	GetProductByID(ctx context.Context, id int64) (*entity.Product, error)
	GetProductsByCategoryID(ctx context.Context, categoryID int64) ([]entity.Product, error)
	UpdateProduct(ctx context.Context, name, description string, price float32, stock, categoryId int64, image string, id int64) error
	DeleteProductByID(ctx context.Context, id int64) error
	//permissions roles interfaces
	GetAllPermissionsRoles(ctx context.Context) ([]entity.PermissionRoles, error)
	GetAllRoles(ctx context.Context) ([]entity.Role, error)
}

type repo struct {
	db *sql.DB
}

func New(_db *sql.DB) Repository {
	return &repo{
		db: _db,
	}
}
