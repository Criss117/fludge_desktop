package dtos

type SignupDTO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Pin      string `json:"pin"`
}
