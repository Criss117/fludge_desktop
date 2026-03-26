package appstate

import (
	iamAggregates "desktop/internal/platform/iam/domain/aggregates"
	"time"
)

type OnStateChangeType string

const (
	SignUp             OnStateChangeType = "SIGN_UP"
	SignIn             OnStateChangeType = "SIGN_IN"
	SignOut            OnStateChangeType = "SIGN_OUT"
	SwitchOrganization OnStateChangeType = "SWITCH_ORGANIZATION"
)

type StateChangeEvent struct {
	Type         OnStateChangeType
	Operator     *iamAggregates.Operator
	Organization *iamAggregates.Organization
}

type SessionState struct {
	ActiveOrganization *iamAggregates.Organization
	ActiveOperator     *ActiveOperator
	UpdatedAt          time.Time
}

type ActiveOperator struct {
	*iamAggregates.Operator
	*iamAggregates.Member
	Teams []*iamAggregates.Team
}

func BuildSessionState(
	operator *iamAggregates.Operator,
	org *iamAggregates.Organization,
) *SessionState {
	if operator == nil {
		return &SessionState{
			ActiveOrganization: nil,
			ActiveOperator:     nil,
			UpdatedAt:          time.Now(),
		}
	}

	var member *iamAggregates.Member
	var teams []*iamAggregates.Team

	if org != nil {
		member = org.FindMemberByOperatorId(operator.ID)
		teams = org.FindTeamsByOperatorId(operator.ID)

		if member == nil {
			return &SessionState{
				ActiveOrganization: nil,
				ActiveOperator:     nil,
				UpdatedAt:          time.Now(),
			}
		}

		if operator.IsRoot() {
			teams = org.Teams
		}
	}

	return &SessionState{
		ActiveOrganization: org,
		ActiveOperator: &ActiveOperator{
			Operator: operator,
			Member:   member,
			Teams:    teams,
		},
		UpdatedAt: time.Now(),
	}
}

func (s *SessionState) IsAuthenticated() bool {
	return s.ActiveOperator != nil
}

func (s *SessionState) HasActiveOrganization() bool {
	return s.ActiveOrganization != nil
}

func (s *SessionState) SetActiveOrganization(org *iamAggregates.Organization) {
	if s.ActiveOperator == nil {
		return
	}

	s.ActiveOrganization = org
	s.UpdatedAt = time.Now()

	if s.ActiveOrganization == nil {
		return
	}

	member := org.FindMemberByOperatorId(s.ActiveOperator.Operator.ID)
	teams := org.FindTeamsByOperatorId(s.ActiveOperator.Operator.ID)

	if member == nil {
		return
	}

	s.ActiveOperator = &ActiveOperator{
		Operator: s.ActiveOperator.Operator,
		Member:   member,
		Teams:    teams,
	}
}

func (s *SessionState) SetActiveOperator(operator *iamAggregates.Operator) {
	s.ActiveOperator = &ActiveOperator{
		Operator: operator,
		Member:   nil,
		Teams:    nil,
	}
	s.UpdatedAt = time.Now()
}

func (s *SessionState) Clear() {
	s.ActiveOrganization = nil
	s.ActiveOperator = nil
	s.UpdatedAt = time.Now()
}

func (s *SessionState) ToAppState() *iamAggregates.AppState {
	var activeOrganizationID *string = nil
	var activeOperatorID *string = nil

	if s.ActiveOrganization != nil {
		activeOrganizationID = &s.ActiveOrganization.ID
	}

	if s.ActiveOperator != nil {
		activeOperatorID = &s.ActiveOperator.Operator.ID
	}

	return iamAggregates.NewAppState(
		activeOrganizationID,
		activeOperatorID,
		s.UpdatedAt,
	)
}

// func (s *SessionState) Can(permission Permission) bool
