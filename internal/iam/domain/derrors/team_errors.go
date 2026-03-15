package derrors

import "errors"

var (
	ErrTeamMemberNotFound              = errors.New("El miembro del equipo no se encuentra")
	ErrTeamMemberAlreadyExists         = errors.New("El miembro del equipo ya existe")
	ErrOperatorNotMemberOfOrganization = errors.New("El operador no es miembro de la organización")
	ErrTeamNotFound                    = errors.New("El equipo no se encuentra")
	ErrTeamPermissionsEmpty            = errors.New("El equipo no tiene permisos")
	ErrPermissionListEmpty             = errors.New("La lista de permisos no puede estar vacía")
)
