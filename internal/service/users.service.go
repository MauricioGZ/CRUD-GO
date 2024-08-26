package service

import (
	"context"
	"errors"

	"github.com/MauricioGZ/CRUD-GO/encryption"
	"github.com/MauricioGZ/CRUD-GO/internal/model"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *serv) RegisterUser(ctx context.Context, firstName, lastName, email, password string) error {
	user, err := s.repo.GetUserByEmail(ctx, email)

	if user != nil {
		if err != nil {
			return err
		}
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	encryptedPassword := encryption.ToBase64(bb)

	err = s.repo.SaveUser(
		ctx,
		firstName,
		lastName,
		email,
		encryptedPassword,
	)

	return err
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(user.Password)
	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return &model.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}
