package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/MauricioGZ/CRUD-GO/internal/entity"
)

const (
	qryInsertUser = `	insert into USERS (
															firstName,
															lastName,
															email,
															password,
															roleId,
															createdAt
															)
										values (?,?,?,?,?,?);`
	qryGetUserByEmail = `	select 
													id, 
													firstName, 
													lastName, 
													email,
													password,
													createdAt
												from USERS
												where email = ?;`
)

const (
	AdminRole    int64 = 1
	SellerRole   int64 = 2
	CustomerRole int64 = 3
)

func (r *repo) SaveUser(ctx context.Context, firstName, lastName, email, password string) error {
	_, err := r.db.ExecContext(
		ctx,
		qryInsertUser,
		firstName,
		lastName,
		email,
		password,
		CustomerRole,
		time.Now().UTC(),
	)
	fmt.Println(err)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRowContext(
		ctx,
		qryGetUserByEmail,
		email,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
