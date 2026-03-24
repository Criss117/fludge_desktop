package valueobjects

import (
	"errors"
	"regexp"
)

type Pin struct {
	pin string
}

var (
	ErrPinEmpty = errors.New("El campo 'pin' no puede estar vacío")
	ErrPinValid = errors.New("El campo 'pin' no es válido")
)

func NewPin(pin string) (Pin, error) {
	if pin == "" {
		return Pin{}, ErrPinEmpty
	}

	if !IsValidPin(pin) {
		return Pin{}, ErrPinValid
	}

	return Pin{
		pin: pin,
	}, nil
}

func ReconstitutePin(pin string) Pin {
	return Pin{
		pin: pin,
	}
}

func IsValidPin(pin string) bool {
	// Esta regex es un balance entre simplicidad y efectividad.
	// Verifica: longitud mínima de 6 caracteres, número y letra.
	pattern := `^[a-zA-Z0-9]{6,}$`

	match, _ := regexp.MatchString(pattern, pin)
	return match
}

func (p Pin) Value() string {
	return p.pin
}

func (p Pin) Equals(other Pin) bool {
	return p.pin == other.pin
}

func (p Pin) ValidatePin(pin string) bool {
	return p.pin == pin
}
