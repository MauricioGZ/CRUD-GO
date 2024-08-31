package entity

import "time"

type Payment struct {
	ID            int64
	OrderID       int64
	PaymentMethod string
	Amount        float32
	PaymentDate   time.Time
	Status        string
}
