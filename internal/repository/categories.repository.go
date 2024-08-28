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
	_, err := r.db.ExecContext(
		ctx,
		qryInsertCategory,
		name,
		description,
		parentID,
	)

	return err
}

func (r *repo) GetCategoryByID(ctx context.Context, id int64) (*entity.Categories, error) {
	var category entity.Categories
	err := r.db.QueryRowContext(
		ctx,
		qryGetCategoryByID,
		id,
	).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.ParentID,
	)

	//TODO: check the case if no rows is retrieved
	if err != nil {
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
