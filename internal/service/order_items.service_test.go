package service

import (
	"context"
	"testing"
)

func TestRegisterOrderItem(t *testing.T) {
	testCases := []struct {
		Name          string
		OrderID       int64
		ProductID     int64
		Quantity      int64
		Price         float32
		ExpectedError error
	}{
		{
			Name:          "RegisterOrderItem: Success",
			OrderID:       validOrderID,
			ProductID:     validProductID,
			Quantity:      100,
			Price:         100.0,
			ExpectedError: nil,
		},
		{
			Name:          "RegisterOrderItem: Invalid Order ID",
			OrderID:       invalidOrderID,
			ProductID:     validProductID,
			Quantity:      100,
			Price:         100.0,
			ExpectedError: ErrOrderDoesNotExist,
		},
		{
			Name:          "RegisterOrderItem: Invalid Product ID",
			OrderID:       validOrderID,
			ProductID:     invalidProductID,
			Quantity:      100,
			Price:         100.0,
			ExpectedError: ErrProductDoesNotExist,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.RegisterOrderItem(ctx, tc.OrderID, tc.ProductID, tc.Quantity, tc.Price)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestGetOrderItemsByUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		ExpectedError error
	}{
		{
			Name:          "GetOrderItemsByUser: Success",
			Email:         existingUserEmail,
			ExpectedError: nil,
		},
		{
			Name:          "GetOrderItemsByUser: User does not exist",
			Email:         unexistingUserEmail,
			ExpectedError: ErrUserDoesntExist,
		},
		{
			Name:          "GetOrderItemsByUser: User does not have orders",
			Email:         existingUserEmailWithoutOrders,
			ExpectedError: ErrNoOrdersRegistered,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.GetOrderItemsByUser(ctx, tc.Email)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
