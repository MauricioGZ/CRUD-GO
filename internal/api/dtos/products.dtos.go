package dtos

type RegisterProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       int64   `json:"stock"`
	CategoryID  int64   `json:"category_id"`
	Image       string  `json:"image"`
}
