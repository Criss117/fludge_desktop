package appstate

import (
	"desktop/internal/platform/iam/application/responses"
)

type ActiveOperatorResponse struct {
	*responses.OperatorResponse `json:"operator"`
	*responses.MemberResponse   `json:"member"`
	Teams                       []*responses.TeamResponse `json:"teams"`
}

type SessionStateResponse struct {
	ActiveOrganization *responses.OrganizationResponse `json:"activeOrganization"`
	ActiveOperator     *ActiveOperatorResponse         `json:"activeOperator"`
}

func SessionStateResponseFromDomain(sessionState *SessionState) *SessionStateResponse {
	if sessionState == nil {
		return nil
	}

	var activeOrganization *responses.OrganizationResponse = nil
	var activeOperator *ActiveOperatorResponse = nil

	if sessionState.ActiveOrganization != nil {
		activeOrganization = responses.OrganizationResponseFromDomain(sessionState.ActiveOrganization)
	}

	if sessionState.ActiveOperator != nil {
		activeOperator = &ActiveOperatorResponse{
			OperatorResponse: responses.OperatorResponseFromDomain(sessionState.ActiveOperator.Operator),
			MemberResponse:   responses.MemberResponseFromDomain(sessionState.ActiveOperator.Member),
			Teams:            make([]*responses.TeamResponse, len(sessionState.ActiveOperator.Teams)),
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
