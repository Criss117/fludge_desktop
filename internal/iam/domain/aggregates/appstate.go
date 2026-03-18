package aggregates

import (
	"desktop/internal/iam/domain/derrors"
	"time"
)

type PrimitiveAppState struct {
	ActiveOrganization *PrimitiveOrganization
	ActiveOperator     *PrimitiveOperator
	UpdatedAt          time.Time
}

type AppState struct {
	activeOrganization *Organization
	activeOperator     *Operator
	updatedAt          time.Time
}

func NewAppState(activeOrganization *Organization, operator *Operator, updatedAt time.Time) AppState {
	return AppState{
		activeOrganization: activeOrganization,
		activeOperator:     operator,
		updatedAt:          updatedAt,
	}
}

func (as *AppState) SwitchOrganization(organization *Organization) error {
	if as.activeOperator == nil {
		return derrors.ErrNoActiveOperator
	}

	if as.activeOrganization == nil {
		return derrors.ErrAppStateNotSetOrganization
	}

	if organization.FindMemberByOperatorId(as.activeOperator.id) == nil {
		return derrors.ErrAppStateNotSetOrganization
	}

	as.activeOrganization = organization
	as.updatedAt = time.Now()

	return nil
}

func (as *AppState) SetActiveOperator(operator *Operator) {
	as.activeOrganization = nil
	as.activeOperator = operator
	as.updatedAt = time.Now()
}

func (as *AppState) ActiveOrganization() *Organization {
	return as.activeOrganization
}

func (as *AppState) ActiveOperator() *Operator {
	return as.activeOperator
}

func (as *AppState) ToValues() PrimitiveAppState {
	var activeOperator *PrimitiveOperator = nil
	var activeOrganization *PrimitiveOrganization = nil

	if as.activeOrganization != nil {
		primitiveOrg := as.activeOrganization.ToValues()
		activeOrganization = &primitiveOrg
	}

	if as.activeOperator != nil {
		primitiveOperator := as.activeOperator.ToValues()
		activeOperator = &primitiveOperator
	}

	return PrimitiveAppState{
		ActiveOrganization: activeOrganization,
		ActiveOperator:     activeOperator,
		UpdatedAt:          as.updatedAt,
	}
}
