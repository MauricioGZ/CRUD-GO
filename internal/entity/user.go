package entity

import "time"

type User struct {
	ID        int64     `db:"id"`
	FirstName string    `db:"firstName"`
	LastName  string    `db:"lastName"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"createdAt"`
}
