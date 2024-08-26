package model

type Address struct {
	ID          int64  `json:"id"`
	AddressType string `json:"address_type"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	ZipCode     string `json:"zip_code"`
}
