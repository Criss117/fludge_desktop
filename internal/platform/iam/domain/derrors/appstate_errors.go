package derrors

import "errors"

var (
	ErrAppStateNotFound             = errors.New("El estado de la aplicación no se encuentra")
	ErrAppStateNotUpdated           = errors.New("No se ha actualizado el estado de la aplicación")
	ErrNoActiveOperator             = errors.New("No hay operador activo")
	ErrNoActiveOrganization         = errors.New("No hay organización activa")
	ErrAppStateNotSetOperator       = errors.New("No se ha establecido el operador activo")
	ErrAppStateNotSetOrganization   = errors.New("No se ha establecido la organización activa")
	ErrAppStateOrganizationMismatch = errors.New("La organización activa no coincide con la solicitada")
)
