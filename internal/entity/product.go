package entity

import "time"

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       float32
	Stock       int64
	CategoryID  int64
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
