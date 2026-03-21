package derrors

import "errors"

var (
	ErrCategoryNameTooShort      = errors.New("El nombre de la categoria es muy corto")
	ErrCategoryNameAlreadyExists = errors.New("El nombre de la categoria ya existe")
)
