package service

import (
	"context"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		FirstName     string
		LastName      string
		Email         string
		Password      string
		Role          string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser: Success",
			FirstName:     "Validname",
			LastName:      "Validsurname",
			Email:         unexistingUserEmail,
			Password:      validPassword,
			Role:          "Customer",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser: Already exists",
			FirstName:     "Validname",
			LastName:      "Validsurname",
			Email:         existingUserEmail,
			Password:      validPassword,
			Role:          "Customer",
			ExpectedError: ErrUserAlreadyExists,
		},
		{
			Name:          "RegisterUser: Invalid role",
			FirstName:     "Validname",
			LastName:      "Validsurname",
			Email:         unexistingUserEmail,
			Password:      validPassword,
			Role:          "",
			ExpectedError: ErrRolesNotInitialized,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.RegisterUser(ctx, tc.FirstName, tc.LastName, tc.Email, tc.Password, tc.Role)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "LoginUser: Success",
			Email:         existingUserEmail,
			Password:      validPassword,
			ExpectedError: nil,
		},
		{
			Name:          "LoginUser: Invalid password",
			Email:         existingUserEmail,
			Password:      invalidPassword,
			ExpectedError: ErrInvalidCredentials,
		},
		{
			Name:          "LoginUser: User does not exist",
			Email:         unexistingUserEmail,
			Password:      invalidPassword,
			ExpectedError: ErrInvalidCredentials,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.LoginUser(ctx, tc.Email, tc.Password)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
