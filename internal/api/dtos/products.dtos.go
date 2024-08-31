package dtos

type RegisterProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       int64   `json:"stock"`
	CategoryID  int64   `json:"category_id"`
	Image       string  `json:"image"`
}

type GetProductByID struct {
	ID int64 `param:"id"`
}

type GetProductByCategory struct {
	CategoryName string `query:"categoryName"`
}

type UpdateProductByID struct {
	ID          int64   `param:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       int64   `json:"stock"`
	CategoryID  int64   `json:"category_id"`
	Image       string  `json:"image"`
}

type DeleteProductByID struct {
	ID int64 `param:"id"`
}
