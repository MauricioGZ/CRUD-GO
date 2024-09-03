package service

import (
	"context"
	"testing"
)

func TestRegisterAddress(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		AddressType   string
		Address       string
		City          string
		State         string
		Country       string
		ZipCode       string
		ExpectedError error
	}{
		{
			Name:          "RegisterAddress: Success",
			Email:         existingUserEmail,
			AddressType:   "billing",
			Address:       "valid street",
			City:          "valid city",
			State:         "valid state",
			Country:       "valid country",
			ZipCode:       "0000001",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterAddress: Success",
			Email:         unexistingUserEmail,
			AddressType:   "billing",
			Address:       "valid street",
			City:          "valid city",
			State:         "valid state",
			Country:       "valid country",
			ZipCode:       "0000001",
			ExpectedError: ErrUserDoesntExist,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.RegisterAddress(
				ctx,
				tc.Email,
				tc.AddressType,
				tc.Address,
				tc.City,
				tc.State,
				tc.Country,
				tc.ZipCode,
			)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestGetAllAddresses(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		ExpectedError error
	}{
		{
			Name:          "RegisterAddress: Success",
			Email:         existingUserEmail,
			ExpectedError: nil,
		},
		{
			Name:          "RegisterAddress: Success",
			Email:         unexistingUserEmail,
			ExpectedError: ErrUserDoesntExist,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.GetAllAddresses(ctx, tc.Email)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
