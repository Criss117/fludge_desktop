package aggregates

import (
	"desktop/internal/iam/domain/derrors"
	"desktop/internal/iam/domain/valueobjects"
	"desktop/internal/shared/lib"
	"time"
)

type PrimitiveTeam struct {
	ID             string
	Name           string
	OrganizationID string
	Permissions    []string
	Description    *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	Members        []*PrimitiveTeamMember
}

type Team struct {
	id             string
	name           string
	organizationID string
	permissions    []valueobjects.Permission
	description    *string
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      *time.Time
	members        []*TeamMember
}

func DefaultTeam(organizationID string) *Team {
	description := "Este equipo es el administrador de la organización"

	return &Team{
		id:             lib.GenerateUUID(),
		name:           "Administradores",
		organizationID: organizationID,
		permissions:    valueobjects.AllPermissions(),
		description:    &description,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
		deletedAt:      nil,
		members:        nil,
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
		id:             lib.GenerateUUID(),
		name:           name,
		organizationID: organizationID,
		permissions:    validPermissions,
		description:    description,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
		deletedAt:      nil,
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
		id:             id,
		name:           name,
		organizationID: organizationID,
		permissions:    permissions,
		description:    description,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
		deletedAt:      deletedAt,
		members:        members,
	}
}

func (t *Team) Delete() {
	now := time.Now()
	t.deletedAt = &now
	t.updatedAt = now
}

func (t *Team) IsActive() bool {
	if t.deletedAt != nil {
		return false
	}

	return true
}

func (t *Team) ID() string {
	return t.id
}

func (t *Team) Members() []*TeamMember {
	return t.members
}

func (t *Team) OperatorIsMember(operatorId string) bool {
	for _, member := range t.members {
		if member.OperatorID() == operatorId {
			return true
		}
	}
	return false
}

func (t *Team) Equals(other *Team) bool {
	return t.id == other.id
}

func (t *Team) ToValues() PrimitiveTeam {
	permissions := make([]string, len(t.permissions))

	for i, p := range t.permissions {
		permissions[i] = p.Value()
	}

	members := make([]*PrimitiveTeamMember, len(t.members))

	for i, member := range t.members {
		tMember := member.ToValues()

		members[i] = &tMember
	}

	return PrimitiveTeam{
		ID:             t.id,
		Name:           t.name,
		OrganizationID: t.organizationID,
		Permissions:    permissions,
		Description:    t.description,
		CreatedAt:      t.createdAt,
		UpdatedAt:      t.updatedAt,
		DeletedAt:      t.deletedAt,
		Members:        members,
	}
}

func (t *Team) FindMemberByOperatorId(operatorId string) *TeamMember {
	for _, member := range t.members {
		if member.OperatorID() == operatorId {
			return member
		}
	}
	return nil
}

func (t *Team) AddMember(member *TeamMember) error {
	if t.FindMemberByOperatorId(member.OperatorID()) != nil {
		return derrors.ErrTeamMemberAlreadyExists
	}

	t.members = append(t.members, member)
	t.updatedAt = time.Now()

	return nil
}

func (t *Team) RemoveMember(member *TeamMember) error {
	for i, m := range t.members {
		if m.Equals(member) {
			t.members = append(t.members[:i], t.members[i+1:]...)
			t.updatedAt = time.Now()
			return nil
		}
	}

	return derrors.ErrTeamMemberNotFound
}

func (t *Team) CountMembers() int {
	return len(t.members)
}
