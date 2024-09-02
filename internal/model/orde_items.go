package model

import "time"

type OrderItem struct {
	ID        int64   `json:"id"`
	ProductID int64   `json:"product_id"`
	Quantity  int64   `json:"quantity"`
	Price     float32 `json:"price"`
}

type OrderItemsByUser struct {
	OrderID    int64       `json:"order_id"`
	OrderDate  time.Time   `json:"order_date"`
	Status     string      `json:"status"`
	TotalPrice float32     `json:"total_price"`
	OrderItems []OrderItem `json:"order_items"`
}
