package derrors

import "errors"

var (
	ErrOrganizationNotFound      = errors.New("El organización no se encuentra")
	ErrOrganizationAlreadyExists = errors.New("El organización ya existe")
)
