package service

import (
	"context"
	"testing"
)

func TestRegisterOrder(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Status        string
		totalPrice    float32
		ExpectedError error
	}{
		{
			Name:          "RegisterOrder: Success",
			Email:         existingUserEmail,
			Status:        "pending",
			totalPrice:    100,
			ExpectedError: nil,
		},
		{
			Name:          "RegisterOrder: User does not exist",
			Email:         unexistingUserEmail,
			Status:        "pending",
			totalPrice:    100,
			ExpectedError: ErrUserDoesntExist,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.RegisterOrder(ctx, tc.Email, tc.Status, tc.totalPrice)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestGetOrderByID(t *testing.T) {
	testCases := []struct {
		Name          string
		ID            int64
		Role          string
		ExpectedError error
	}{
		{
			Name:          "GetOrderByID: Success",
			ID:            validOrderID,
			Role:          "Seller",
			ExpectedError: nil,
		},
		{
			Name:          "GetOrderByID: Invalid User Role",
			ID:            validOrderID,
			Role:          "Customer",
			ExpectedError: ErrInvalidPermissions,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.GetOrderByID(ctx, tc.ID, tc.Role)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestGetOrdersByUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		ExpectedError error
	}{
		{
			Name:          "GetOrdersByUser: Success",
			Email:         existingUserEmail,
			ExpectedError: nil,
		},
		{
			Name:          "GetOrdersByUser: User does not exist",
			Email:         unexistingUserEmail,
			ExpectedError: ErrUserDoesntExist,
		},
		{
			Name:          "GetOrdersByUser: Success",
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

			_, err := s.GetOrdersByUser(ctx, tc.Email)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
