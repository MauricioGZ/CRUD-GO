package entity

type Address struct {
	ID          int64
	UserID      int64
	AddressType string
	Address     string
	City        string
	State       string
	Country     string
	ZipCode     string
}
