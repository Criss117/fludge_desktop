package valueobjects

import (
	"errors"
	"regexp"
)

type Email struct {
	email string
}

var (
	ErrEmailEmpty = errors.New("El campo 'email' no puede estar vacío")
	ErrEmailValid = errors.New("El campo 'email' no es válido")
)

func NewEmail(email string) (Email, error) {
	if email == "" {
		return Email{}, ErrEmailEmpty
	}

	if !IsValidEmail(email) {
		return Email{}, ErrEmailValid
	}

	return Email{
		email: email,
	}, nil
}

func ReconstituteEmail(email string) Email {
	return Email{
		email: email,
	}
}

func IsValidEmail(email string) bool {
	// Esta regex es un balance entre simplicidad y efectividad.
	// Verifica: usuario + @ + dominio + extensión de mínimo 2 caracteres.
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, _ := regexp.MatchString(pattern, email)
	return match
}

func (e *Email) Value() string {
	return e.email
}
