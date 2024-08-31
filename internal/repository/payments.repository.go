package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/MauricioGZ/CRUD-GO/internal/entity"
)

const (
	qryInsertPayment = `insert into PAYMENTS(
												orderId,
												paymenMethod,
												amount,
												paymentDate,
												status
											)
											values (?,?,?,?,?);`
	qryGetPaymentByID = `	select
													id,
													orderId,
													paymentMethod,
													amount,
													paymentDate,
													status
												froms PAYMENTS
												where id = ?;`
	qryGetPaymentByOrderID = `	select
																id,
																orderId,
																paymentMethod,
																amount,
																paymentDate,
																status
															froms PAYMENTS
															where orderId = ?;`
)

func (r *repo) InsertPayment(ctx context.Context, orderId int64, paymentMethod string, amount float32, paymentDate time.Time, status string) error {
	_, err := r.db.ExecContext(
		ctx,
		qryInsertPayment,
		orderId,
		paymentMethod,
		amount,
		paymentDate,
		status,
	)

	return err
}

func (r *repo) GetPaymentByID(ctx context.Context, id int64) (*entity.Payment, error) {
	var payment entity.Payment
	err := r.db.QueryRowContext(
		ctx,
		qryGetPaymentByID,
		id,
	).Scan(
		&payment.ID,
		&payment.OrderID,
		&payment.PaymentMethod,
		&payment.Amount,
		&payment.PaymentDate,
		&payment.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &payment, nil
}

func (r *repo) GetPaymentByOrderID(ctx context.Context, orderID int64) (*entity.Payment, error) {
	var payment entity.Payment
	err := r.db.QueryRowContext(
		ctx,
		qryGetPaymentByOrderID,
		orderID,
	).Scan(
		&payment.ID,
		&payment.OrderID,
		&payment.PaymentMethod,
		&payment.Amount,
		&payment.PaymentDate,
		&payment.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &payment, nil
}
