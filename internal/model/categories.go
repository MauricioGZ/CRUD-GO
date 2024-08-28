package model

type Categories struct {
	ID            int64        `json:"id"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	ChildCategory []Categories `json:"child_category"`
}
