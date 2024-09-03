package service

import (
	"context"
	"errors"

	"github.com/MauricioGZ/CRUD-GO/internal/model"
)

var (
	ErrOrderDoesNotExist  = errors.New("order does not exist")
	ErrNoOrdersRegistered = errors.New("no orders registered")
)

func (s *serv) RegisterOrder(ctx context.Context, email string, status string, totalPrice float32) (*model.OrderResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)

	if user == nil {
		if err != nil {
			return nil, err
		}
		return nil, ErrUserDoesntExist
	}

	orderID, err := s.repo.InsertOrder(ctx, user.ID, status, totalPrice)

	if err != nil {
		return nil, err
	}

	return &model.OrderResponse{ID: *orderID}, nil
}

func (s *serv) GetOrderByID(ctx context.Context, id int64, role string) (*model.Order, error) {
	//only sellers and admins may access to the orders by its id
	if !mayUpdate(role) {
		return nil, ErrInvalidPermissions
	}

	o, err := s.repo.GetOrderByID(ctx, id)

	if err != nil {
		return nil, err
	}

	order := model.Order{
		ID:         o.ID,
		OrderDate:  o.OrderDate,
		Status:     o.Status,
		TotalPrice: o.TotalPrice,
	}

	return &order, nil
}

func (s *serv) GetOrdersByUser(ctx context.Context, email string) ([]model.Order, error) {
	var orders []model.Order
	user, err := s.repo.GetUserByEmail(ctx, email)
	if user == nil {
		if err != nil {
			return nil, err
		}
		return nil, ErrUserDoesntExist
	}

	oo, err := s.repo.GetOrdersByUserID(ctx, user.ID)

	if oo == nil {
		if err != nil {
			return nil, err
		}
		return nil, ErrNoOrdersRegistered
	}

	for _, o := range oo {
		orders = append(orders, model.Order{
			ID:         o.ID,
			OrderDate:  o.OrderDate,
			Status:     o.Status,
			TotalPrice: o.TotalPrice,
		})
	}

	return orders, nil
}
