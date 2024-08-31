package repository

import (
	"context"
	"database/sql"
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
													roleId,
													createdAt
												from USERS
												where email = ?;`
	qryDeleteUserByEmail = `delete from USERS
													where email = ?;`
)

func (r *repo) InsertUser(ctx context.Context, firstName, lastName, email, password string, roleID int64) error {
	_, err := r.db.ExecContext(
		ctx,
		qryInsertUser,
		firstName,
		lastName,
		email,
		password,
		roleID,
		time.Now().UTC(),
	)

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
		&user.RoleID,
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

func (r *repo) DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := r.db.ExecContext(ctx, qryDeleteUserByEmail, email)
	return err
}
