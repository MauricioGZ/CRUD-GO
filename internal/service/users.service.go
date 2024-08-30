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

func (s *serv) RegisterUser(ctx context.Context, firstName, lastName, email, password, role string) error {
	var roleID int64
	user, err := s.repo.GetUserByEmail(ctx, email)

	if user != nil {
		if err != nil {
			return err
		}
		return ErrUserAlreadyExists
	}

	roleID = getRoleID(role)
	if roleID == 0 {
		return ErrRolesNotInitialized
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	encryptedPassword := encryption.ToBase64(bb)

	err = s.repo.InsertUser(
		ctx,
		firstName,
		lastName,
		email,
		encryptedPassword,
		roleID,
	)

	return err
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if user == nil {
		if err != nil {
			return nil, err
		}
		return nil, ErrInvalidCredentials
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

	role := getRole(user.RoleID)

	if role == "" {
		return nil, ErrRolesNotInitialized
	}

	return &model.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      role,
	}, nil
}
