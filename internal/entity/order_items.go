package entity

import "time"

type OrderItem struct {
	ID        int64
	OrderID   int64
	ProductID int64
	Quantity  int64
	Price     float32
}

type OrderItemByUserID struct {
	OrderItem
	OrderDate  time.Time
	Status     string
	TotalPrice float32
}
