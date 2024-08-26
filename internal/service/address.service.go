package service

import (
	"context"

	"github.com/MauricioGZ/CRUD-GO/internal/model"
)

func (s *serv) RegisterAddress(ctx context.Context, email, addressType, address, city, state, country, zipCode string) error {
	u, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return err
	}

	err = s.repo.SaveAddress(ctx, u.ID, addressType, address, city, state, country, zipCode)

	return err
}

func (s *serv) UpdateAddress(ctx context.Context, id int64, addressType, address, city, state, country, zipCode string) error {

	err := s.repo.UpdateAddressByID(ctx, id, addressType, address, city, state, country, zipCode)

	return err
}

func (s *serv) GetAllAddresses(ctx context.Context, email string) ([]model.Address, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	aa, err := s.repo.GetAddressesByUserId(ctx, u.ID)

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
