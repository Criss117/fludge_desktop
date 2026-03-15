package aggregates

import (
	"desktop/internal/shared/lib"
	"time"
)

type PrimitiveTeamMember struct {
	ID             string
	TeamID         string
	OperatorID     string
	OrganizationID string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

type TeamMember struct {
	id             string
	teamID         string
	operatorID     string
	organizationID string
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      *time.Time
}

func NewTeamMember(teamID, operatorID, organizationID string) *TeamMember {
	return &TeamMember{
		id:             lib.GenerateUUID(),
		teamID:         teamID,
		operatorID:     operatorID,
		organizationID: organizationID,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
		deletedAt:      nil,
	}
}

func ReconstituteTeamMember(id, teamID, operatorID, organizationID string, createdAt, updatedAt time.Time, deletedAt *time.Time) *TeamMember {
	return &TeamMember{
		id:             id,
		teamID:         teamID,
		operatorID:     operatorID,
		organizationID: organizationID,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
		deletedAt:      deletedAt,
	}
}

func (tm *TeamMember) Delete() {
	now := time.Now()
	tm.deletedAt = &now
	tm.updatedAt = now
}

func (tm *TeamMember) IsActive() bool {
	if tm.deletedAt != nil {
		return false
	}

	return true
}

func (tm *TeamMember) ID() string {
	return tm.id
}

func (tm *TeamMember) OperatorID() string {
	return tm.operatorID
}

func (tm *TeamMember) Equals(other *TeamMember) bool {
	return tm.id == other.id
}

func (tm *TeamMember) ToValues() PrimitiveTeamMember {
	return PrimitiveTeamMember{
		ID:             tm.id,
		TeamID:         tm.teamID,
		OperatorID:     tm.operatorID,
		OrganizationID: tm.organizationID,
		CreatedAt:      tm.createdAt,
		UpdatedAt:      tm.updatedAt,
		DeletedAt:      tm.deletedAt,
	}
}
