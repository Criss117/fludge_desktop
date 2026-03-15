package aggregates

import (
	"desktop/internal/iam/domain/derrors"
	"time"
)

type AppState struct {
	activeOrganizationId *string
	activeOperator       *Operator
	updatedAt            time.Time
}

func NewAppState(activeOrganizationId *string, operator *Operator, updatedAt time.Time) AppState {
	return AppState{
		activeOrganizationId: activeOrganizationId,
		activeOperator:       operator,
		updatedAt:            updatedAt,
	}
}

func (as *AppState) SetOrganization(organization *Organization) error {
	if as.activeOperator == nil {
		return derrors.ErrNoActiveOperator
	}

	if as.activeOrganizationId == nil {
		return derrors.ErrAppStateNotSetOrganization
	}

	if organization.FindMemberByOperatorId(*as.activeOrganizationId) == nil {
		return derrors.ErrAppStateNotSetOrganization
	}

	as.activeOrganizationId = &organization.id
	as.updatedAt = time.Now()

	return nil
}

func (as *AppState) SetActiveOperator(operator *Operator) {
	as.activeOperator = operator
	as.updatedAt = time.Now()
}
