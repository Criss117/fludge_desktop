package derrors

import "errors"

var (
	ErrOperatorNotFound                = errors.New("No se encontro el operador")
	ErrOperatorAlreadyExists           = errors.New("El operador ya existe")
	ErrOperatorAlreadyExistsByEmail    = errors.New("El operador ya existe con el mismo email")
	ErrOperatorAlreadyExistsByUsername = errors.New("El operador ya existe con el mismo nombre de usuario")
	ErrInvalidCredentials              = errors.New("Credenciales inválidas")
)
