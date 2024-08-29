package service

import (
	"context"
	"errors"
)

var (
	ErrCategoryDoesNotExist = errors.New("category does not exist")
)

func (s *serv) RegisterProduct(ctx context.Context, name, description string, price float32, stock, categoryId int64, image string) error {
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
