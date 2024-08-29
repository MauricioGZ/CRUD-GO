package entity

import "time"

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
	RoleID    int64
	CreatedAt time.Time
}
