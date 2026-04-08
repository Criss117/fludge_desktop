package aggregates

import (
	"desktop/internal/platform/iam/domain/derrors"
	"desktop/internal/platform/iam/domain/valueobjects"
	"desktop/internal/shared/lib"
	"time"
)

type Team struct {
	ID             string
	Name           string
	OrganizationID string
	Permissions    valueobjects.PermissionList
	Description    *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Members        []*TeamMember
}

func DefaultTeam(organizationID string) *Team {
	description := "Este equipo es el administrador de la organización"

	return &Team{
		ID:             lib.GenerateUUID(),
		Name:           "Administradores",
		OrganizationID: organizationID,
		Permissions:    valueobjects.AllPermissions(),
		Description:    &description,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      nil,
		Members:        nil,
	}
}

func NewTeam(
	name, organizationID string,
	permissions []string,
	description *string,
) (*Team, error) {
	if len(permissions) == 0 {
		return nil, derrors.ErrPermissionListEmpty
	}

	validPermissions, errPermissions := valueobjects.NewPermissionList(permissions)

	if errPermissions != nil {
		return nil, errPermissions
	}

	return &Team{
		ID:             lib.GenerateUUID(),
		Name:           name,
		OrganizationID: organizationID,
		Permissions:    validPermissions,
		Description:    description,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      nil,
	}, nil
}

func ReconstituteTeam(
	id, name, organizationID string,
	permissions valueobjects.PermissionList,
	description *string,
	createdAt, updatedAt time.Time,
	deletedAt *time.Time,
	members []*TeamMember,
) *Team {
	return &Team{
		ID:             id,
		Name:           name,
		OrganizationID: organizationID,
		Permissions:    permissions,
		Description:    description,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		DeletedAt:      deletedAt,
		Members:        members,
	}
}

func (t *Team) Delete() {
	now := time.Now()
	t.DeletedAt = &now
	t.UpdatedAt = now
}

func (t *Team) IsActive() bool {
	if t.DeletedAt != nil {
		return false
	}

	return true
}

func (t *Team) OperatorIsMember(operatorId string) bool {
	for _, member := range t.Members {
		if member.OperatorID == operatorId {
			return true
		}
	}
	return false
}

func (t *Team) Equals(other *Team) bool {
	return t.ID == other.ID
}

func (t *Team) FindMemberByOperatorId(operatorId string) *TeamMember {
	for _, member := range t.Members {
		if member.OperatorID == operatorId {
			return member
		}
	}
	return nil
}

func (t *Team) AddMember(member *TeamMember) error {
	if t.FindMemberByOperatorId(member.OperatorID) != nil {
		return derrors.ErrTeamMemberAlreadyExists
	}

	t.Members = append(t.Members, member)
	t.UpdatedAt = time.Now()

	return nil
}

func (t *Team) RemoveMember(member *TeamMember) error {
	for i, m := range t.Members {
		if m.Equals(member) {
			t.Members = append(t.Members[:i], t.Members[i+1:]...)
			t.UpdatedAt = time.Now()
			return nil
		}
	}

	return derrors.ErrTeamMemberNotFound
}

func (t *Team) CountMembers() int {
	return len(t.Members)
}

func (t *Team) Update(
	name string,
	description *string,
	permissions []string,
) error {
	if len(permissions) == 0 {
		return derrors.ErrPermissionListEmpty
	}

	validPermissions, errPermissions := valueobjects.NewPermissionList(permissions)

	if errPermissions != nil {
		return errPermissions
	}

	t.Name = name
	t.Permissions = validPermissions
	t.Description = description
	t.UpdatedAt = time.Now()

	return nil
}
