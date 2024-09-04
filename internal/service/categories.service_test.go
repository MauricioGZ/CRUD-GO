package service

import (
	"context"
	"testing"
)

func TestRegisterCategory(t *testing.T) {
	testCases := []struct {
		Name          string
		Role          string
		CategoryName  string
		Description   string
		ParentID      int64
		ExpectedError error
	}{
		{
			Name:          "RegisterCategory: Success",
			Role:          "Admin",
			CategoryName:  newCategory,
			Description:   "Some description",
			ParentID:      1,
			ExpectedError: nil,
		},
		{
			Name:          "RegisterCategory: Invalid Role",
			Role:          "Customer",
			CategoryName:  existingCategory,
			Description:   "Some description",
			ParentID:      1,
			ExpectedError: ErrInvalidPermissions,
		},
		{
			Name:          "RegisterCategory: Category dupplicated",
			Role:          "Admin",
			CategoryName:  existingCategory,
			Description:   "Some description",
			ParentID:      1,
			ExpectedError: ErrCategoryAlreadyExists,
		},
		{
			Name:          "RegisterCategory: Parent id does not exist",
			Role:          "Admin",
			CategoryName:  newCategory,
			Description:   "Some description",
			ParentID:      100,
			ExpectedError: ErrParentCategoryDoesNotExist,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.RegisterCategory(ctx, tc.Role, tc.CategoryName, tc.Description, tc.ParentID)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
