package derrors

import "errors"

var (
	ErrOrganizationNotFound      = errors.New("El organización no se encuentra")
	ErrOrganizationAlreadyExists = errors.New("El organización ya existe")
	ErrRootMemberCannotBeAdded   = errors.New("No se puede agregar un miembro root")
	ErrOperatorIsNotMemberOfOrg  = errors.New("El operador no es miembro de la organización")
	ErrTeamAlreadyExists         = errors.New("El equipo ya existe")
)
