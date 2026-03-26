package appstate

import (
	"desktop/internal/platform/iam/application/responses"
)

type ActiveOperatorResponse struct {
	*responses.Operator `json:"operator"`
	*responses.Member   `json:"member"`
	Teams               []*responses.Team `json:"teams"`
}

type SessionStateResponse struct {
	ActiveOrganization *responses.Organization `json:"activeOrganization"`
	ActiveOperator     *ActiveOperatorResponse `json:"activeOperator"`
}

func SessionStateResponseFromDomain(sessionState *SessionState) SessionStateResponse {
	if sessionState == nil {
		return SessionStateResponse{}
	}

	var activeOrganization *responses.Organization = nil
	var activeOperator *ActiveOperatorResponse = nil

	if sessionState.ActiveOrganization != nil {
		ao := responses.OrganizationFromDomain(sessionState.ActiveOrganization)

		activeOrganization = &ao
	}

	if sessionState.ActiveOperator != nil {
		op := responses.OperatorFromDomain(sessionState.ActiveOperator.Operator)

		activeOperator = &ActiveOperatorResponse{
			Operator: &op,
		}
	}

	return SessionStateResponse{
		ActiveOrganization: activeOrganization,
		ActiveOperator:     activeOperator,
	}
}
