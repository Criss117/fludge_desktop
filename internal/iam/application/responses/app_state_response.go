package responses

import (
	"desktop/internal/iam/domain/aggregates"
)

type ResponseAppState struct {
	ActiveOrganization *OrganizationResponse `json:"activeOrganization"`
	ActiveOperator     *OperatorResponse     `json:"activeOperator"`
}

func ResponseAppStateFromDomain(
	appState *aggregates.AppState,
) *ResponseAppState {

	var activeOrganization *OrganizationResponse = nil
	var activeOperator *OperatorResponse = nil

	if appState.ActiveOrganization != nil {
		activeOrganization = OrganizationResponseFromDomain(appState.ActiveOrganization)
	}

	if appState.ActiveOperator != nil {
		activeOperator = OperatorResponseFromDomain(appState.ActiveOperator)
	}

	return &ResponseAppState{
		ActiveOrganization: activeOrganization,
		ActiveOperator:     activeOperator,
	}
}
