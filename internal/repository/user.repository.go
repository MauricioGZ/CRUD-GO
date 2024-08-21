package repository

import (
	"context"
	"time"

	"github.com/MauricioGZ/CRUD-GO/internal/entity"
)

const (
	qryInsertUser = `	insert into users (
															firstName,
															lastName,
															email,
															password,
															createdAt
															)
										values (?,?,?,?,?);`
	qryGetUserByEmail = `	select 
													id, 
													firstName, 
													lastName, 
													email,
													password,
													createdAt
												from users
												where email = ?;`
)

func (r *repo) SaveUser(ctx context.Context, firstName, lastName, email, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser,
		firstName,
		lastName,
		email,
		password,
		time.Now().UTC())
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
		return nil, err
	}

	return &user, nil
}
