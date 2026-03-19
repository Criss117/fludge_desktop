package aggregates

import (
	"desktop/internal/iam/domain/valueobjects"
	"desktop/internal/shared/lib"
	"time"
)

type Member struct {
	ID             string
	OrganizationID string
	OperatorID     string
	Role           valueobjects.MemberRole
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

func NewMember(organizationID, operatorID string, role valueobjects.MemberRole) *Member {
	return &Member{
		ID:             lib.GenerateUUID(),
		OrganizationID: organizationID,
		OperatorID:     operatorID,
		Role:           role,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      nil,
	}
}

func ReconstituteMember(id, organizationID, operatorID string, role string, createdAt, updatedAt time.Time, deletedAt *time.Time) *Member {
	return &Member{
		ID:             id,
		OrganizationID: organizationID,
		OperatorID:     operatorID,
		Role:           valueobjects.ReconstituteMemberRole(role),
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		DeletedAt:      deletedAt,
	}
}

func (m *Member) Delete() {
	now := time.Now()
	m.DeletedAt = &now
	m.UpdatedAt = now
}

func (m *Member) IsRoot() bool {
	return m.Role.IsRoot()
}

func (m *Member) IsActive() bool {
	if m.DeletedAt != nil {
		return false
	}

	return true
}

func (m *Member) Equals(other *Member) bool {
	return m.ID == other.ID
}
