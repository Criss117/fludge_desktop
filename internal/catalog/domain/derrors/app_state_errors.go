package derrors

import "errors"

var (
	ErrAppStateNotFound     = errors.New("Error: El estado de la aplicación no se encuentra")
	ErrNoActiveOrganization = errors.New("Error: No hay organización activa")
)
