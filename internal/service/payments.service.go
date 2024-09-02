package service

import (
	"context"
	"time"
)

func (s *serv) RegisterPayment(ctx context.Context, orderID int64, paymentMethod string, amount float32, paymentDate time.Time, status string) error {
	order, err := s.repo.GetOrderByID(ctx, orderID)

	if order == nil {
		if err != nil {
			return err
		}
		return ErrOrderDoesNotExist
	}

	err = s.repo.InsertPayment(
		ctx,
		orderID,
		paymentMethod,
		amount,
		paymentDate,
		status,
	)

	return err
}
