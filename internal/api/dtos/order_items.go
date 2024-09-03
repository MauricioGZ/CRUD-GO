package dtos

type OrderItem struct {
	ProductID int64   `json:"product_id"`
	Quantity  int64   `json:"quantity"`
	Price     float32 `json:"price"`
}
