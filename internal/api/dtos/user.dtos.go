package dtos

type RegisterUser struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type LoginUser struct {
	Email    string
	Password string
}
