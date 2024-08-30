package service

import (
	"context"
	"errors"

	"github.com/MauricioGZ/CRUD-GO/internal/model"
)

var (
	ErrInvalidPermissions  = errors.New("invalid permissions")
	ErrProductDoesNotExist = errors.New("product does not exist")
)

func (s *serv) RegisterProduct(ctx context.Context, role, name, description string, price float32, stock, categoryId int64, image string) error {
	if !mayCreate(role) {
		return ErrInvalidPermissions
	}

	category, err := s.repo.GetCategoryByID(ctx, categoryId)
	if category == nil {
		if err != nil {
			return err
		}
		return ErrCategoryDoesNotExist
	}

	err = s.repo.InsertProduct(
		ctx,
		name,
		description,
		price,
		stock,
		categoryId,
		image,
	)

	return err
}

func (s *serv) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product

	pp, err := s.repo.GetAllProducts(ctx)

	if err != nil {
		return nil, err
	}

	for _, p := range pp {
		products = append(
			products,
			model.Product{
				ID:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Stock:       p.Stock,
				CategoryID:  p.CategoryID,
				Image:       p.Image,
				CreatedAt:   p.CreatedAt,
				UpdatedAt:   p.UpdatedAt,
			},
		)
	}

	return products, nil
}

func (s *serv) GetProductByID(ctx context.Context, id int64) (*model.Product, error) {
	p, err := s.repo.GetProductByID(ctx, id)

	if p == nil {
		if err != nil {
			return nil, err
		}
		return nil, ErrProductDoesNotExist
	}

	if err != nil {
		return nil, err
	}

	product := model.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		CategoryID:  p.CategoryID,
		Image:       p.Image,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}

	return &product, nil
}

func (s *serv) GetProductsByCategory(ctx context.Context, categoryName string) ([]model.Product, error) {
	var products []model.Product
	category, err := s.repo.GetCategoryByName(ctx, categoryName)

	if category == nil {
		if err != nil {
			return nil, err
		}
		return nil, ErrCategoryDoesNotExist
	}

	pp, err := s.repo.GetProductsByCategoryID(ctx, category.ID)

	if err != nil {
		return nil, err
	}

	for _, p := range pp {
		products = append(
			products,
			model.Product{
				ID:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Stock:       p.Stock,
				CategoryID:  p.CategoryID,
				Image:       p.Image,
				CreatedAt:   p.CreatedAt,
				UpdatedAt:   p.UpdatedAt,
			},
		)
	}

	return products, nil
}

func (s *serv) UpdateProduct(ctx context.Context, role, name, description string, price float32, stock, categoryId int64, image string, id int64) error {
	if !mayUpdate(role) {
		return ErrInvalidPermissions
	}

	p, err := s.repo.GetProductByID(ctx, id)

	if p == nil {
		if err != nil {
			return err
		}
		return ErrProductDoesNotExist
	}

	err = s.repo.UpdateProduct(
		ctx,
		name,
		description,
		price,
		stock,
		categoryId,
		image,
		id,
	)

	return err
}

func (s *serv) DeleteProductByID(ctx context.Context, role string, id int64) error {
	if !mayDelete(role) {
		return ErrInvalidPermissions
	}

	p, err := s.repo.GetProductByID(ctx, id)

	if p == nil {
		if err != nil {
			return err
		}
		return ErrProductDoesNotExist
	}

	err = s.repo.DeleteProductByID(ctx, id)

	return err
}
