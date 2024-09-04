package service

import (
	"context"
	"errors"

	"github.com/MauricioGZ/CRUD-GO/internal/entity"
	"github.com/MauricioGZ/CRUD-GO/internal/model"
)

var (
	ErrCategoryAlreadyExists      = errors.New("category already exists")
	ErrParentCategoryDoesNotExist = errors.New("parent category does not exist")
	ErrCategoryDoesNotExist       = errors.New("category does not exist")
)

func (s *serv) RegisterCategory(ctx context.Context, role, name, description string, parentID int64) error {
	if role != "Admin" {
		return ErrInvalidPermissions
	}

	if parentID != 0 {
		parentCategory, err := s.repo.GetCategoryByID(ctx, parentID)
		if parentCategory == nil {
			if err != nil {
				return err
			}
			return ErrParentCategoryDoesNotExist
		}
	}

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

func (s *serv) GetAllCategories(ctx context.Context) ([]model.Categories, error) {
	var categories []model.Categories
	var category model.Categories
	cc, err := s.repo.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	//cc is sorted since the db
	for i, c := range cc {
		//only append the categories without parent category
		if c.ParentID == 0 {
			category = model.Categories{
				ID:          c.ID,
				Name:        c.Name,
				Description: c.Description,
			}
			appendChildCategory(&category, cc, i)
			categories = append(categories, category)
		}
	}

	return categories, nil
}

func appendChildCategory(parent *model.Categories, cc []entity.Categories, lastIndex int) {
	var childCategory model.Categories
	for i := lastIndex + 1; i < len(cc); i++ {
		//only append the child categories if the parent id matchs
		if parent.ID == cc[i].ParentID {
			childCategory = model.Categories{
				ID:          cc[i].ID,
				Name:        cc[i].Name,
				Description: cc[i].Description,
			}
			//check for nested categories
			appendChildCategory(&childCategory, cc, i)
			parent.ChildCategory = append(parent.ChildCategory, childCategory)
		}
	}
}
