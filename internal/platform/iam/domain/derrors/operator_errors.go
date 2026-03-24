package derrors

import "errors"

var (
	ErrOperatorNotFound                  = errors.New("No se encontro el operador")
	ErrOperatorAlreadyExists             = errors.New("El operador ya existe")
	ErrOperatorEmailAlreadyExists        = errors.New("El operador ya existe con el mismo email")
	ErrOperatorUsernameAlreadyExists     = errors.New("El operador ya existe con el mismo nombre de usuario")
	ErrInvalidCredentials                = errors.New("Credenciales inválidas")
	ErrOperatorCanBeMemberInOrganization = errors.New("El operador debe ser miembro de una organización")
	ErrOperatorMustBeRoot                = errors.New("El operador debe ser root")
	ErrOperatorTypeInvalid               = errors.New("Tipo de operador inválido")
)
