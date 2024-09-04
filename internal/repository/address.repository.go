package repository

import (
	"context"
	"database/sql"

	"github.com/MauricioGZ/CRUD-GO/internal/entity"
)

const (
	qryInsertAddress = `insert into ADDRESSES(
												userId,
												addressType,
												address,
												city,
												state,
												country,
												zipCode)
											values (?,?,?,?,?,?,?);`
	qryGetAddressesByUserID = `	select 
																id,
																userId,
																addressType,
																address,
																city,
																state,
																country,
																zipCode
															from ADDRESSES
															where userId = ?;`
	qryDeleteAddressByID = `delete
													from ADDRESSES
													where id = ? AND userId = ?;`
	qryUpdateAddressByID = `update ADDRESSES
													set
														addressType = ?,
														address = ?,
														city = ?,
														state = ?,
														country = ?,
														zipCode = ?
													where id = ? AND userId = ?;`
	qryGetAddressByID = `	select 
												id,
												userId,
												addressType,
												address,
												city,
												state,
												country,
												zipCode
											from ADDRESSES
											where id = ? AND userId = ?;`
)

func (r *repo) SaveAddress(ctx context.Context, userId int64, addressType, address, city, state, country, zipCode string) error {
	_, err := r.db.ExecContext(
		ctx,
		qryInsertAddress,
		userId,
		addressType,
		address,
		city,
		state,
		country,
		zipCode,
	)

	return err
}

func (r *repo) GetAddressesByUserId(ctx context.Context, userId int64) ([]entity.Address, error) {
	var address entity.Address
	var addresses []entity.Address

	rows, err := r.db.QueryContext(
		ctx,
		qryGetAddressesByUserID,
		userId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&address.ID,
			&address.UserID,
			&address.AddressType,
			&address.Address,
			&address.City,
			&address.State,
			&address.Country,
			&address.ZipCode,
		)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (r *repo) DeleteAddressByID(ctx context.Context, id, userID int64) error {
	_, err := r.db.ExecContext(
		ctx,
		qryDeleteAddressByID,
		id,
		userID,
	)
	return err
}

func (r *repo) UpdateAddressByID(ctx context.Context, id, userID int64, addressType, address, city, state, country, zipCode string) error {
	_, err := r.db.ExecContext(
		ctx,
		qryUpdateAddressByID,
		addressType,
		address,
		city,
		state,
		country,
		zipCode,
		id,
		userID,
	)

	return err
}

func (r *repo) GetAddressByID(ctx context.Context, id, userID int64) (*entity.Address, error) {
	var address entity.Address
	err := r.db.QueryRowContext(
		ctx,
		qryGetAddressByID,
		id,
		userID,
	).Scan(
		&address.ID,
		&address.UserID,
		&address.AddressType,
		&address.Address,
		&address.City,
		&address.State,
		&address.Country,
		&address.ZipCode,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &address, nil
}
