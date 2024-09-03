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
			CategoryID:    validCategoryID,
			Image:         "image",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterProduct: Invalid permissions",
			Role:          "Customer",
			Description:   "Some description",
			Price:         100.0,
			Stock:         100,
			CategoryID:    validCategoryID,
			Image:         "image",
			ExpectedError: ErrInvalidPermissions,
		},
		{
			Name:          "RegisterProduct: Invalid category",
			Role:          "Admin",
			Description:   "Some description",
			Price:         100.0,
			Stock:         100,
			CategoryID:    invalidCategoryID,
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
			ID:            validProductID,
			ExpectedError: nil,
		},
		{
			Name:          "GetProductByID: Product doest not exist",
			ID:            invalidProductID,
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
			CategoryName:  existingCategory,
			ExpectedError: nil,
		},
		{
			Name:          "GetProductsByCategory: Invalid category",
			CategoryName:  unexistingCategory,
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
			CategoryID:    validCategoryID,
			Image:         "some image",
			ID:            validProductID,
			ExpectedError: nil,
		},
		{
			Name:          "UpdateProduct: Invalid permissions",
			Role:          "Customer",
			ProductName:   "some name",
			Description:   "Some description",
			Price:         100.0,
			Stock:         100,
			CategoryID:    validCategoryID,
			Image:         "some image",
			ID:            validProductID,
			ExpectedError: ErrInvalidPermissions,
		},
		{
			Name:          "UpdateProduct: Invalid product id",
			Role:          "Admin",
			ProductName:   "some name",
			Description:   "Some description",
			Price:         100.0,
			Stock:         100,
			CategoryID:    validCategoryID,
			Image:         "some image",
			ID:            invalidProductID,
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
			ID:            validProductID,
			ExpectedError: nil,
		},
		{
			Name:          "DeleteProductByID: Invalid permissions",
			Role:          "Customer",
			ID:            validProductID,
			ExpectedError: ErrInvalidPermissions,
		},
		{
			Name:          "DeleteProductByID: Invalid product id",
			Role:          "Admin",
			ID:            invalidProductID,
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
