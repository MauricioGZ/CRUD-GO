package service

import (
	"context"

	"github.com/MauricioGZ/CRUD-GO/internal/model"
)

func (s *serv) RegisterOrderItem(ctx context.Context, orderID, productID, quantity int64, price float32) error {
	order, err := s.repo.GetOrderByID(ctx, orderID)

	if order == nil {
		if err != nil {
			return err
		}
		return ErrOrderDoesNotExist
	}

	product, err := s.repo.GetProductByID(ctx, productID)

	if product == nil {
		if err != nil {
			return err
		}
		return ErrProductDoesNotExist
	}

	err = s.repo.InsertOrderItem(ctx, orderID, productID, quantity, price)

	return err
}

func (s *serv) GetOrderItemsByUser(ctx context.Context, email string) ([]model.OrderItemsByUser, error) {
	//help variable to count the indexes of each order
	var index int64 = 0
	var orders []model.OrderItemsByUser
	//map to gather the indexes of each order
	ordersIndex := make(map[int64]int64)
	user, err := s.repo.GetUserByEmail(ctx, email)

	if user == nil {
		if err != nil {
			return nil, err
		}
		return nil, ErrUserDoesntExist
	}

	ooii, err := s.repo.GetOrderItemsByUserID(ctx, user.ID)

	if ooii == nil {
		if err != nil {
			return nil, err
		}
		return nil, ErrNoOrdersRegistered
	}

	for _, oi := range ooii {
		_, ok := ordersIndex[oi.OrderID]
		//append orders only by different order ids
		if !ok {
			ordersIndex[oi.OrderID] = index
			index++
			orders = append(orders, model.OrderItemsByUser{
				OrderID:    oi.OrderID,
				OrderDate:  oi.OrderDate,
				Status:     oi.Status,
				TotalPrice: oi.TotalPrice,
			})
		}

		orders[ordersIndex[oi.OrderID]].OrderItems = append(orders[ordersIndex[oi.OrderID]].OrderItems, model.OrderItem{
			ID:        oi.ID,
			ProductID: oi.ProductID,
			Quantity:  oi.Quantity,
			Price:     oi.Price,
		})
	}

	return orders, nil
}
