package aggregates

import (
	"desktop/internal/iam/domain/derrors"
	"time"
)

type AppState struct {
	ActiveOrganization *Organization
	ActiveOperator     *Operator
	UpdatedAt          time.Time
}

func NewAppState(activeOrganization *Organization, operator *Operator, updatedAt time.Time) AppState {
	return AppState{
		ActiveOrganization: activeOrganization,
		ActiveOperator:     operator,
		UpdatedAt:          updatedAt,
	}
}

func (as *AppState) SwitchOrganization(organization *Organization) error {
	if as.ActiveOperator == nil {
		return derrors.ErrNoActiveOperator
	}

	if organization.FindMemberByOperatorId(as.ActiveOperator.ID) == nil {
		return derrors.ErrAppStateNotSetOrganization
	}

	as.ActiveOrganization = organization
	as.UpdatedAt = time.Now()

	return nil
}

func (as *AppState) SetActiveOperator(operator *Operator) {
	as.ActiveOrganization = nil
	as.ActiveOperator = operator
	as.UpdatedAt = time.Now()
}
