package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/MauricioGZ/CRUD-GO/internal/entity"
)

const (
	qryInsertOrder = `insert into ORDERS(
											userId,
											orderDate,
											status,
											totalPrice
										)
										values (?,?,?,?);`
	qryGetOrderByID = ` select
												id,
												userId,
												orderDate,
												status,
												totalPrice
											from ORDERS
											where id = ?;`
	qryGetOrderByUserID = ` select
														id,
														userId,
														orderDate,
														status,
														totalPrice
													from ORDERS
													where userId = ?;`
)

func (r *repo) InsertOrder(ctx context.Context, userID int64, status string, totalPrice float32) error {
	_, err := r.db.ExecContext(
		ctx,
		qryInsertOrder,
		userID,
		time.Now().UTC(),
		status,
		totalPrice,
	)

	return err
}

func (r *repo) GetOrderByID(ctx context.Context, id int64) (*entity.Order, error) {
	var order entity.Order
	err := r.db.QueryRowContext(
		ctx,
		qryGetOrderByID,
		id,
	).Scan(
		&order.ID,
		&order.UserID,
		&order.OrderDate,
		&order.Status,
		&order.TotalPrice,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}

func (r *repo) GetOrdersByUserID(ctx context.Context, userID int64) ([]entity.Order, error) {
	var order entity.Order
	var orders []entity.Order
	rows, err := r.db.QueryContext(
		ctx,
		qryGetOrderByUserID,
		userID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&order.ID,
			&order.UserID,
			&order.OrderDate,
			&order.Status,
			&order.TotalPrice,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
