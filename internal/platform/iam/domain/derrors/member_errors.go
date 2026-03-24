package derrors

import "errors"

var (
	ErrMemberNotFound      = errors.New("El miembro no se encuentra")
	ErrMemberRoleInvalid   = errors.New("El rol de miembro no es valido")
	ErrMemberAlreadyExists = errors.New("El miembro ya existe")
)
