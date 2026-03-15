package appstate

import iamAggregates "desktop/internal/iam/domain/aggregates"

func BuildSessionState(
	activeOrganization *iamAggregates.Organization,
	activeOperator *iamAggregates.Operator,
	operatorMember *iamAggregates.Member,
	operatorTeams []*iamAggregates.Team,
) *SessionState {

	return &SessionState{
		ActiveOrganization: activeOrganization,
		ActiveOperator: &ActiveOperator{
			Operator: activeOperator,
			Member:   operatorMember,
			Teams:    operatorTeams,
		},
	}
}
