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

func SessionStateResponseFromDomain(sessionState *SessionState) *SessionStateResponse {
	if sessionState == nil {
		return nil
	}

	var activeOrganization *responses.Organization = nil
	var activeOperator *ActiveOperatorResponse = nil

	if sessionState.ActiveOrganization != nil {
		activeOrganization = responses.OrganizationResponseFromDomain(sessionState.ActiveOrganization)
	}

	if sessionState.ActiveOperator != nil {
		activeOperator = &ActiveOperatorResponse{
			Operator: responses.OperatorResponseFromDomain(sessionState.ActiveOperator.Operator),
			Member:   responses.MemberResponseFromDomain(sessionState.ActiveOperator.Member),
			Teams:    make([]*responses.Team, len(sessionState.ActiveOperator.Teams)),
		}

		for i, team := range sessionState.ActiveOperator.Teams {
			activeOperator.Teams[i] = responses.TeamResponseFromDomain(team)
		}
	}

	return &SessionStateResponse{
		ActiveOrganization: activeOrganization,
		ActiveOperator:     activeOperator,
	}
}
