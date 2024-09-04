package service

import (
	"context"
	"errors"

	"github.com/MauricioGZ/CRUD-GO/internal/model"
)

var (
	ErrUserDoesntExist       = errors.New("user does not exist")
	ErrAddressDoesNotExist   = errors.New("address does not exist")
	ErrNoAddressesRegistered = errors.New("no addresses registered")
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

func (s *serv) UpdateAddress(ctx context.Context, id int64, email, addressType, address, city, state, country, zipCode string) error {
	user, err := s.repo.GetUserByEmail(ctx, email)

	if user == nil {
		if err != nil {
			return err
		}
		return ErrUserDoesntExist
	}

	a, err := s.repo.GetAddressByID(ctx, id, user.ID)

	if a == nil {
		if err != nil {
			return err
		}
		return ErrAddressDoesNotExist
	}

	if addressType != "billing" && addressType != "shipping" {
		addressType = a.AddressType
	}
	if address == "" {
		address = a.Address
	}
	if city == "" {
		city = a.City
	}
	if state == "" {
		state = a.State
	}
	if country == "" {
		country = a.Country
	}
	if zipCode == "" {
		zipCode = a.ZipCode
	}

	err = s.repo.UpdateAddressByID(ctx, id, user.ID, addressType, address, city, state, country, zipCode)

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
	if aa == nil {
		if err != nil {
			return nil, err
		}
		return nil, ErrNoAddressesRegistered
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

func (s *serv) DeleteAddress(ctx context.Context, id int64, email string) error {
	user, err := s.repo.GetUserByEmail(ctx, email)

	if user == nil {
		if err != nil {
			return err
		}
		return ErrUserDoesntExist
	}

	a, err := s.repo.GetAddressByID(ctx, id, user.ID)

	if a == nil {
		if err != nil {
			return err
		}
		return ErrAddressDoesNotExist
	}

	err = s.repo.DeleteAddressByID(ctx, id, user.ID)

	return err
}
