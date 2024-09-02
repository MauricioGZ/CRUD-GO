package repository

import (
	"context"
	"database/sql"

	"github.com/MauricioGZ/CRUD-GO/internal/entity"
)

const (
	qryInsertOrderItem = `insert into ORDER_ITEMS(
													orderId,
													productId,
													quantity,
													price
												)
												values (?,?,?,?);`
	qryGetOrderItemsByOrderID = `	select
																	id,
																	orderId,
																	productId,
																	quantity,
																	price
																from ORDER_ITEMS
																where orderId = ?;`
	qryUpdateOrderItemByID = `update ORDER_ITEMS
														set
															productId = ?,
															quantity = ?,
															price = ?
														where orderId = ?;`
	qryDeleteOrderItemsByOrderID = `delete 
																	from ORDER_ITEMS
																	where orderId = ?;`
	qryDeleteOrderItemByID = `	delete 
															from ORDER_ITEMS
															where id = ?;`
	qryGetOrderItemsByUserID = `select
																ORDERS.id,
																ORDERS.orderDate,
																ORDERS.status,
																ORDERS.totalPrice,
																ORDER_ITEMS.id,
																ORDER_ITEMS.productId,
																ORDER_ITEMS.quantity,
																ORDER_ITEMS.price
															from ORDERS
															join ORDER_ITEMS on ORDERS.id = ORDER_ITEMS.orderId
															where ORDERS.userId = ?;`
)

func (r *repo) InsertOrderItem(ctx context.Context, orderID, productID, quantity int64, price float32) error {
	_, err := r.db.ExecContext(
		ctx,
		qryInsertOrderItem,
		orderID,
		productID,
		quantity,
		price,
	)

	return err
}

func (r *repo) GetOrderItemsByOrderId(ctx context.Context, orderID int64) ([]entity.OrderItem, error) {
	var orderItem entity.OrderItem
	var orderItems []entity.OrderItem

	rows, err := r.db.QueryContext(
		ctx,
		qryGetOrderItemsByOrderID,
		orderID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&orderItem.ID,
			&orderItem.OrderID,
			&orderItem.ProductID,
			&orderItem.Quantity,
			&orderItem.Price,
		)

		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}

func (r *repo) DeleteOrderItemsByOrderID(ctx context.Context, orderID int64) error {
	_, err := r.db.ExecContext(
		ctx,
		qryDeleteOrderItemsByOrderID,
		orderID,
	)
	return err
}

func (r *repo) DeleteOrderItemByID(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(
		ctx,
		qryDeleteOrderItemByID,
		id,
	)
	return err
}

func (r *repo) GetOrderItemsByUserID(ctx context.Context, userID int64) ([]entity.OrderItemByUserID, error) {
	var orderItem entity.OrderItemByUserID
	var orderItems []entity.OrderItemByUserID

	rows, err := r.db.QueryContext(
		ctx,
		qryGetOrderItemsByUserID,
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
			&orderItem.OrderID,
			&orderItem.OrderDate,
			&orderItem.Status,
			&orderItem.TotalPrice,
			&orderItem.ID,
			&orderItem.ProductID,
			&orderItem.Quantity,
			&orderItem.Price,
		)

		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}
