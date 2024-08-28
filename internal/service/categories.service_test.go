package service

import (
	"context"
	"testing"
)

func TestRegisterCategory(t *testing.T) {
	testCases := []struct {
		Name          string
		CategoryName  string
		Description   string
		ParentID      int64
		ExpectedError error
	}{
		{
			Name:          "RegisterCategory: Success",
			CategoryName:  "New Category",
			Description:   "Some description",
			ParentID:      1,
			ExpectedError: nil,
		},
		{
			Name:          "RegisterCategory: Category dupplicated",
			CategoryName:  "Existing Category",
			Description:   "Some description",
			ParentID:      1,
			ExpectedError: ErrCategoryAlreadyExists,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.RegisterCategory(ctx, tc.CategoryName, tc.Description, tc.ParentID)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
