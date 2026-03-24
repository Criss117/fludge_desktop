package aggregates

import (
	"time"
)

type AppState struct {
	ActiveOrganizationID *string
	ActiveOperatorID     *string
	UpdatedAt            time.Time
}

func NewAppState(activeOrganizationID, activeOperatorId *string, updatedAt time.Time) *AppState {
	return &AppState{
		ActiveOrganizationID: activeOrganizationID,
		ActiveOperatorID:     activeOperatorId,
		UpdatedAt:            updatedAt,
	}
}

func (as *AppState) SetActiveOrganization(organizationId string) error {
	as.ActiveOrganizationID = &organizationId
	as.UpdatedAt = time.Now()

	return nil
}

func (as *AppState) SetActiveOperator(operatorId *string) {
	as.ActiveOrganizationID = nil
	as.ActiveOperatorID = operatorId
	as.UpdatedAt = time.Now()
}
