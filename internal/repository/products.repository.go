package repository

import (
	"context"
	"time"
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
