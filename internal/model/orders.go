package model

import "time"

type Order struct {
	ID         int64     `json:"id"`
	OrderDate  time.Time `json:"order_date"`
	Status     string    `json:"status"`
	TotalPrice float32   `json:"total_price"`
}
