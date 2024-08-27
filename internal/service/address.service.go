package service

import (
	"context"
	"errors"

	"github.com/MauricioGZ/CRUD-GO/internal/model"
)

var (
	ErrUserDoesntExist = errors.New("user doesn't exist")
)

func (s *serv) RegisterAddress(ctx context.Context, email, addressType, address, city, state, country, zipCode string) error {
	user, err := s.repo.GetUserByEmail(ctx, email)

	if user == nil {
		if err != nil {
			return err
		}
		return ErrUserDoesntExist
	}

	err = s.repo.SaveAddress(ctx, user.ID, addressType, address, city, state, country, zipCode)

	return err
}

func (s *serv) UpdateAddress(ctx context.Context, id int64, addressType, address, city, state, country, zipCode string) error {

	err := s.repo.UpdateAddressByID(ctx, id, addressType, address, city, state, country, zipCode)

	return err
}

func (s *serv) GetAllAddresses(ctx context.Context, email string) ([]model.Address, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)

	if user == nil {
		if err != nil {
			return nil, err
		}
		return nil, ErrUserDoesntExist
	}

	aa, err := s.repo.GetAddressesByUserId(ctx, user.ID)

	if err != nil {
		return nil, err
	}

	addresses := []model.Address{}
	for _, a := range aa {
		addresses = append(
			addresses,
			model.Address{
				ID:          a.ID,
				AddressType: a.AddressType,
				Address:     a.Address,
				City:        a.City,
				State:       a.State,
				Country:     a.Country,
				ZipCode:     a.ZipCode,
			})
	}

	return addresses, nil
}

func (s *serv) DeleteAddress(ctx context.Context, id int64) error {
	err := s.repo.DeleteAddressByID(ctx, id)
	return err
}
