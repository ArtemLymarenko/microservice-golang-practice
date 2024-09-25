package dto

type RegisterUser struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email" validation:"required"`
	Password  string `json:"password" validation:"required"`
}
