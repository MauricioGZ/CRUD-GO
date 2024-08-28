package service

import (
	"context"
	"errors"
)

var (
	ErrCategoryAlreadyExists = errors.New("category already exists")
)

func (s *serv) RegisterCategory(ctx context.Context, name, description string, parentID int64) error {
	category, err := s.repo.GetCategoryByName(ctx, name)

	if category != nil {
		if err != nil {
			return err
		}
		return ErrCategoryAlreadyExists
	}

	err = s.repo.InsertCategory(
		ctx,
		name,
		description,
		parentID,
	)

	return err
}
