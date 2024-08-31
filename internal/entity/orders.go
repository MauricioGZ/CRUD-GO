package entity

import "time"

type Order struct {
	ID         int64
	UserID     int64
	OrderDate  time.Time
	Status     string
	TotalPrice float32
}
