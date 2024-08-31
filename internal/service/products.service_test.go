package service

import (
	"context"
	"testing"
)

func TestRegisterProduct(t *testing.T) {
	testCases := []struct {
		Name          string
		Role          string
		Description   string
		Price         float32
		Stock         int64
		CategoryID    int64
		Image         string
		ExpectedError error
	}{
		{
			Name:          "RegisterProduct: Success",
			Role:          "Admin",
			Description:   "Some description",
			Price:         100.0,
			Stock:         100,
			CategoryID:    1,
			Image:         "image",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterProduct: Invalid permissions",
			Role:          "Customer",
			Description:   "Some description",
			Price:         100.0,
			Stock:         100,
			CategoryID:    1,
			Image:         "image",
			ExpectedError: ErrInvalidPermissions,
		},
		{
			Name:          "RegisterProduct: Invalid category",
			Role:          "Admin",
			Description:   "Some description",
			Price:         100.0,
			Stock:         100,
			CategoryID:    100,
			Image:         "image",
			ExpectedError: ErrCategoryDoesNotExist,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
			err := s.RegisterProduct(
				ctx,
				tc.Role,
				tc.Name,
				tc.Description,
				tc.Price,
				tc.Stock,
				tc.CategoryID,
				tc.Image,
			)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestGetProductByID(t *testing.T) {
	testCases := []struct {
		Name          string
		ID            int64
		ExpectedError error
	}{
		{
			Name:          "GetProductByID: Success",
			ID:            1,
			ExpectedError: nil,
		},
		{
			Name:          "GetProductByID: Product doest not exist",
			ID:            100,
			ExpectedError: ErrProductDoesNotExist,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.GetProductByID(ctx, tc.ID)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestGetProductsByCategory(t *testing.T) {
	testCases := []struct {
		Name          string
		CategoryName  string
		ExpectedError error
	}{
		{
			Name:          "GetProductsByCategory: Success",
			CategoryName:  "Existing Category",
			ExpectedError: nil,
		},
		{
			Name:          "GetProductsByCategory: Invalid category",
			CategoryName:  "Unexisting Category",
			ExpectedError: ErrCategoryDoesNotExist,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.GetProductsByCategory(ctx, tc.CategoryName)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	testCases := []struct {
		Name          string
		Role          string
		ProductName   string
		Description   string
		Price         float32
		Stock         int64
		CategoryID    int64
		Image         string
		ID            int64
		ExpectedError error
	}{
		{
			Name:          "UpdateProduct: Success",
			Role:          "Admin",
			ProductName:   "some name",
			Description:   "Some description",
			Price:         100.0,
			Stock:         100,
			CategoryID:    1,
			Image:         "some image",
			ID:            1,
			ExpectedError: nil,
		},
		{
			Name:          "UpdateProduct: Invalid permissions",
			Role:          "Customer",
			ProductName:   "some name",
			Description:   "Some description",
			Price:         100.0,
			Stock:         100,
			CategoryID:    1,
			Image:         "some image",
			ID:            1,
			ExpectedError: ErrInvalidPermissions,
		},
		{
			Name:          "UpdateProduct: Invalid product id",
			Role:          "Admin",
			ProductName:   "some name",
			Description:   "Some description",
			Price:         100.0,
			Stock:         100,
			CategoryID:    2,
			Image:         "some image",
			ID:            100,
			ExpectedError: ErrProductDoesNotExist,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.UpdateProduct(
				ctx,
				tc.Role,
				tc.ProductName,
				tc.Description,
				tc.Price,
				tc.Stock,
				tc.CategoryID,
				tc.Image,
				tc.ID,
			)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})

	}
}

func TestDeleteProductByID(t *testing.T) {
	testCases := []struct {
		Name          string
		Role          string
		ID            int64
		ExpectedError error
	}{
		{
			Name:          "DeleteProductByID: Success",
			Role:          "Admin",
			ID:            1,
			ExpectedError: nil,
		},
		{
			Name:          "DeleteProductByID: Invalid permissions",
			Role:          "Customer",
			ID:            1,
			ExpectedError: ErrInvalidPermissions,
		},
		{
			Name:          "DeleteProductByID: Invalid product id",
			Role:          "Admin",
			ID:            100,
			ExpectedError: ErrProductDoesNotExist,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.DeleteProductByID(ctx, tc.Role, tc.ID)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
