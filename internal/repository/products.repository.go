package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/MauricioGZ/CRUD-GO/internal/entity"
)

const (
	qryInsertProduct = `insert into PRODUCTS(
												name,
												description,
												price,
												stock,
												categoryId,
												image,
												createdAt,
												updatedAt)
											values(?,?,?,?,?,?,?,?);`
	qryGetAllProducts = `	select
													id,
													name,
													description,
													price,
													stock,
													categoryId,
													image,
													createdAt,
													updatedAt
												from PRODUCTS;`
	qryGetProductByID = `	select
													id,
													name,
													description,
													price,
													stock,
													categoryId,
													image,
													createdAt,
													updatedAt
												from PRODUCTS
												where id = ?;`
	qryGetProductsByCategoryID = `select
																	id,
																	name,
																	description,
																	price,
																	stock,
																	categoryId,
																	image,
																	createdAt,
																	updatedAt
																from PRODUCTS
																where categoryId = ?;`
	qryUpdateProductByID = `update PRODUCTS
													set
														name = ?,
														description = ?,
														price = ?,
														stock = ?,
														categoryId = ?,
														image = ?,
														updatedAt = ?
													where id = ?;`
	qryDeleteProductByID = `delete 
													from PRODUCTS
													where id = ?;`
)

func (r *repo) InsertProduct(ctx context.Context, name, description string, price float32, stock, categoryId int64, image string) error {
	_, err := r.db.ExecContext(
		ctx,
		qryInsertProduct,
		name,
		description,
		price,
		stock,
		categoryId,
		image,
		time.Now().UTC(),
		time.Now().UTC(),
	)

	return err
}

func (r *repo) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	var product entity.Product
	var products []entity.Product

	rows, err := r.db.QueryContext(
		ctx,
		qryGetAllProducts,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CategoryID,
			&product.Image,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *repo) GetProductByID(ctx context.Context, id int64) (*entity.Product, error) {
	var product entity.Product

	err := r.db.QueryRowContext(
		ctx,
		qryGetProductByID,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.CategoryID,
		&product.Image,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (r *repo) GetProductsByCategoryID(ctx context.Context, categoryID int64) ([]entity.Product, error) {
	var product entity.Product
	var products []entity.Product

	rows, err := r.db.QueryContext(
		ctx,
		qryGetProductsByCategoryID,
		categoryID,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CategoryID,
			&product.Image,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *repo) UpdateProduct(ctx context.Context, name, description string, price float32, stock, categoryId int64, image string, id int64) error {
	_, err := r.db.ExecContext(
		ctx,
		qryUpdateProductByID,
		name,
		description,
		price,
		stock,
		categoryId,
		image,
		time.Now().UTC(),
		id,
	)

	return err
}

func (r *repo) DeleteProductByID(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(
		ctx,
		qryDeleteProductByID,
		id,
	)

	return err
}
