package derrors

import "errors"

var (
	ErrInvalidDeleteState = errors.New("La operación no se puede realizar porque la no esta marcada como borrada")
)
