package repository

import (
	"context"
	"database/sql"

	"github.com/MauricioGZ/CRUD-GO/internal/entity"
)

const (
	qryInsertCategory = `insert into CATEGORIES(
													name,
													description,
													parentId)
												values(?,?,?);`
	qryInsertCategoryWithoutParent = `insert into CATEGORIES(
																			name,
																			description)
																		values(?,?);`
	qryGetAllCategories = `	select
														id,
														name,
														description,
														parentId
													from CATEGORIES
													ORDER BY parentId ASC;`
	qryGetCategoryByID = `select
													id,
													name,
													description,
													parentId
												from CATEGORIES
												where id = ?;`
	qryGetCategoryByName = `select
													id,
													name,
													description,
													parentId
												from CATEGORIES
												where name = ?;`
	qryUpdateCategoryByID = `	update CATEGORIES
														set
															name = ?,
															descripction = ?,
															parentId = ?
														where id = ?;`
)

func (r *repo) InsertCategory(ctx context.Context, name, description string, parentID int64) error {
	var err error
	if parentID == 0 {
		_, err = r.db.ExecContext(
			ctx,
			qryInsertCategoryWithoutParent,
			name,
			description,
		)
	} else {
		_, err = r.db.ExecContext(
			ctx,
			qryInsertCategory,
			name,
			description,
			parentID,
		)
	}

	return err
}

func (r *repo) GetAllCategories(ctx context.Context) ([]entity.Categories, error) {
	var category entity.Categories
	var categories []entity.Categories
	var parentID sql.NullInt64

	rows, err := r.db.QueryContext(
		ctx,
		qryGetAllCategories,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&parentID,
		)

		if err != nil {
			return nil, err
		}

		if parentID.Valid {
			category.ParentID = parentID.Int64
		} else {
			category.ParentID = 0
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *repo) GetCategoryByID(ctx context.Context, id int64) (*entity.Categories, error) {
	var category entity.Categories
	var parentID sql.NullInt64
	err := r.db.QueryRowContext(
		ctx,
		qryGetCategoryByID,
		id,
	).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&parentID,
	)

	//if parentID is null, that means that the category does not have a parent category
	if parentID.Valid {
		category.ParentID = parentID.Int64
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}

func (r *repo) GetCategoryByName(ctx context.Context, name string) (*entity.Categories, error) {
	var category entity.Categories
	err := r.db.QueryRowContext(
		ctx,
		qryGetCategoryByName,
		name,
	).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.ParentID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}

func (r *repo) UpdateCategoryByID(ctx context.Context, id int64, name, description string, parentID int64) error {
	_, err := r.db.ExecContext(
		ctx,
		qryUpdateCategoryByID,
		name,
		description,
		parentID,
		id,
	)

	return err
}
