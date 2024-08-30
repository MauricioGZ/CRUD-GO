package service

import (
	"context"

	"github.com/MauricioGZ/CRUD-GO/internal/model"
	"github.com/MauricioGZ/CRUD-GO/internal/repository"
)

type Service interface {
	//user services
	RegisterUser(ctx context.Context, firstName, lastName, email, password, role string) error
	LoginUser(ctx context.Context, email, password string) (*model.User, error)
	//address services
	RegisterAddress(ctx context.Context, email, addressType, address, city, state, country, zipCode string) error
	UpdateAddress(ctx context.Context, id int64, addressType, address, city, state, country, zipCode string) error
	GetAllAddresses(ctx context.Context, email string) ([]model.Address, error)
	DeleteAddress(ctx context.Context, id int64) error
	//cateogires services
	RegisterCategory(ctx context.Context, name, description string, parentID int64) error
	GetAllCategories(ctx context.Context) ([]model.Categories, error)
	//products services
	RegisterProduct(ctx context.Context, role, name, description string, price float32, stock, categoryId int64, image string) error
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	GetProductByID(ctx context.Context, id int64) (*model.Product, error)
	GetProductsByCategory(ctx context.Context, categoryName string) ([]model.Product, error)
	UpdateProduct(ctx context.Context, role, name, description string, price float32, stock, categoryId int64, image string, id int64) error
	DeleteProductByID(ctx context.Context, role string, id int64) error
	//permissions roles services
	GetAllPermissionsRoles(ctx context.Context) error
	GetAllRoles(ctx context.Context) error
}

type serv struct {
	repo repository.Repository
}

func New(_repo repository.Repository) Service {
	return &serv{
		repo: _repo,
	}
}
