package aggregates

import (
	"desktop/internal/shared/lib"
	"time"
)

type TeamMember struct {
	ID             string
	TeamID         string
	OperatorID     string
	OrganizationID string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

func NewTeamMember(teamID, operatorID, organizationID string) *TeamMember {
	return &TeamMember{
		ID:             lib.GenerateUUID(),
		TeamID:         teamID,
		OperatorID:     operatorID,
		OrganizationID: organizationID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      nil,
	}
}

func ReconstituteTeamMember(id, teamID, operatorID, organizationID string, createdAt, updatedAt time.Time, deletedAt *time.Time) *TeamMember {
	return &TeamMember{
		ID:             id,
		TeamID:         teamID,
		OperatorID:     operatorID,
		OrganizationID: organizationID,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		DeletedAt:      deletedAt,
	}
}

func (tm *TeamMember) Delete() {
	now := time.Now()
	tm.DeletedAt = &now
	tm.UpdatedAt = now
}

func (tm *TeamMember) IsActive() bool {
	if tm.DeletedAt != nil {
		return false
	}

	return true
}

func (tm *TeamMember) Equals(other *TeamMember) bool {
	return tm.ID == other.ID
}
