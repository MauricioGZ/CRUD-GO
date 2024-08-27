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
		ExpectedError error
	}{
		{
			Name:          "RegisterUser: Success",
			FirstName:     "Validname",
			LastName:      "Validsurname",
			Email:         "email@newuser.com",
			Password:      "validpassword",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser: Already exists",
			FirstName:     "Validname",
			LastName:      "Validsurname",
			Email:         "email@existinguser.com",
			Password:      "validpassword",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.RegisterUser(ctx, tc.FirstName, tc.LastName, tc.Email, tc.Password)
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
			Email:         "email@existinguser.com",
			Password:      "validpassword",
			ExpectedError: nil,
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
