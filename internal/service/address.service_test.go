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

func TestUpdateAddress(t *testing.T) {
	testCases := []struct {
		Name          string
		ID            int64
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
			Name:          "UpdateAddress: Success",
			ID:            validAddresID,
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
			Name:          "UpdateAddress: No Valid User",
			ID:            validAddresID,
			Email:         unexistingUserEmail,
			AddressType:   "billing",
			Address:       "valid street",
			City:          "valid city",
			State:         "valid state",
			Country:       "valid country",
			ZipCode:       "0000001",
			ExpectedError: ErrUserDoesntExist,
		},
		{
			Name:          "UpdateAddress: Valid User Without Addresses",
			ID:            validAddresID,
			Email:         existingUserEmailWithoutAddresses,
			AddressType:   "billing",
			Address:       "valid street",
			City:          "valid city",
			State:         "valid state",
			Country:       "valid country",
			ZipCode:       "0000001",
			ExpectedError: ErrAddressDoesNotExist,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.UpdateAddress(
				ctx,
				tc.ID,
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
			Name:          "RegisterAddress: User does not exist",
			Email:         unexistingUserEmail,
			ExpectedError: ErrUserDoesntExist,
		},
		{
			Name:          "RegisterAddress: User has no addresses",
			Email:         existingUserEmailWithoutAddresses,
			ExpectedError: ErrNoAddressesRegistered,
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

func TestDeleteAddress(t *testing.T) {
	testCases := []struct {
		Name          string
		ID            int64
		Email         string
		ExpectedError error
	}{
		{
			Name:          "DeleteAddress: Success",
			ID:            validAddresID,
			Email:         existingUserEmail,
			ExpectedError: nil,
		},
		{
			Name:          "DeleteAddress: User does not exist",
			ID:            validAddresID,
			Email:         unexistingUserEmail,
			ExpectedError: ErrUserDoesntExist,
		},
		{
			Name:          "DeleteAddress: User has no addresses",
			ID:            validAddresID,
			Email:         existingUserEmailWithoutAddresses,
			ExpectedError: ErrAddressDoesNotExist,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.DeleteAddress(ctx, tc.ID, tc.Email)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
