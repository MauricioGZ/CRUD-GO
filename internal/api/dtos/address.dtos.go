package dtos

type RegisterAddress struct {
	AddressType string `json:"address_type"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	ZipCode     string `json:"zip_code"`
}

type UpdateAddress struct {
	ID int64 `json:"id"`
	RegisterAddress
}
