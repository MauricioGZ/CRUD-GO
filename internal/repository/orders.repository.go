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

func (r *repo) InsertOrder(ctx context.Context, userID int64, orderDate time.Time, status string, totalPrice float64) error {
	_, err := r.db.ExecContext(
		ctx,
		qryInsertOrder,
		userID,
		orderDate,
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

func (r *repo) GetOrderByUserID(ctx context.Context, userID int64) (*entity.Order, error) {
	var order entity.Order
	err := r.db.QueryRowContext(
		ctx,
		qryGetOrderByUserID,
		userID,
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
