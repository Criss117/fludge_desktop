package aggregates

import (
	"desktop/internal/iam/domain/valueobjects"
	"desktop/internal/shared/lib"
	"time"
)

type PrimitiveMember struct {
	ID             string
	OrganizationID string
	OperatorID     string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

type Member struct {
	id             string
	organizationID string
	operatorID     string
	role           valueobjects.MemberRole
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      *time.Time
}

func NewMember(organizationID, operatorID string, role string) (*Member, error) {
	validRole, errRole := valueobjects.NewMemberRole(role)

	if errRole != nil {
		return nil, errRole
	}

	return &Member{
		id:             lib.GenerateUUID(),
		organizationID: organizationID,
		operatorID:     operatorID,
		role:           validRole,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
		deletedAt:      nil,
	}, nil
}

func ReconstituteMember(id, organizationID, operatorID string, role string, createdAt, updatedAt time.Time, deletedAt *time.Time) *Member {
	return &Member{
		id:             id,
		organizationID: organizationID,
		operatorID:     operatorID,
		role:           valueobjects.ReconstituteMemberRole(role),
		createdAt:      createdAt,
		updatedAt:      updatedAt,
		deletedAt:      deletedAt,
	}
}

func (m *Member) Delete() {
	now := time.Now()
	m.deletedAt = &now
	m.updatedAt = now
}

func (m *Member) IsRoot() bool {
	return m.role.IsRoot()
}

func (m *Member) IsActive() bool {
	if m.deletedAt != nil {
		return false
	}

	return true
}

func (m *Member) ID() string {
	return m.id
}

func (m *Member) OperatorID() string {
	return m.operatorID
}

func (m *Member) OrganizationID() string {
	return m.organizationID
}

func (m *Member) Equals(other *Member) bool {
	return m.id == other.id
}

func (m *Member) ToValues() PrimitiveMember {
	return PrimitiveMember{
		ID:             m.id,
		OrganizationID: m.organizationID,
		OperatorID:     m.operatorID,
		Role:           m.role.Value(),
		CreatedAt:      m.createdAt,
		UpdatedAt:      m.updatedAt,
		DeletedAt:      m.deletedAt,
	}
}
