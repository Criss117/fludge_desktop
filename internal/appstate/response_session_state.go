package appstate

import (
	"desktop/internal/iam/application/responses"
	iamAggregates "desktop/internal/iam/domain/aggregates"
)

type ResponseSessionState struct {
	ActiveOrganization *responses.OrganizationResponse `json:"activeOrganization"`
	ActiveOperator     *responses.OperatorResponse     `json:"activeOperator"`
}

func ResponseSessionStateFromDomain(
	activeOrganization *iamAggregates.Organization,
	activeOperator *iamAggregates.Operator,
) *ResponseSessionState {

	operatorResponse := responses.OperatorResponseFromDomain(activeOperator)
	organizationResponse := responses.OrganizationResponseFromDomain(activeOrganization)

	return &ResponseSessionState{
		ActiveOrganization: organizationResponse,
		ActiveOperator:     operatorResponse,
	}
}
