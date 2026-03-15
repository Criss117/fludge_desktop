package appstate

import (
	iamAggregates "desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/valueobjects"
	"errors"
)

var ErrOperatorNotMemberOfOrganization = errors.New("El operador no es miembro de la organizacion")

type ActiveOperator struct {
	*iamAggregates.Operator
	*iamAggregates.Member
	Teams []*iamAggregates.Team
}

type SessionState struct {
	ActiveOrganization *iamAggregates.Organization
	ActiveOperator     *ActiveOperator
}

func NewSessionState(
	activeOrganization *iamAggregates.Organization,
	operator *iamAggregates.Operator,
	member *iamAggregates.Member,
	teams []*iamAggregates.Team,
) (*SessionState, error) {
	operatorIsMember := member.OperatorID() == operator.ID() && member.OrganizationID() == activeOrganization.ID()

	if !operatorIsMember {
		return nil, ErrOperatorNotMemberOfOrganization
	}

	return &SessionState{
		ActiveOrganization: activeOrganization,
		ActiveOperator: &ActiveOperator{
			Operator: operator,
			Member:   member,
			Teams:    teams,
		},
	}, nil
}

func (s *SessionState) IsAuthenticated() bool {
	return s.ActiveOperator != nil
}

func (s *SessionState) HasOrganization() bool {
	return s.ActiveOrganization != nil
}

func (s *SessionState) IsRootOperator() bool {
	return s.ActiveOperator != nil && s.ActiveOperator.Operator.IsRoot()
}

func (s *SessionState) IsOrgRoot() bool {
	return s.ActiveOperator != nil && s.ActiveOperator.Member.IsRoot()
}

func (s *SessionState) Can(permission []valueobjects.Permission) bool {
	if s.IsRootOperator() || s.IsOrgRoot() {
		return true
	}
	// delega al PermissionResolver con los datos ya cargados
	return false
}

func (s *SessionState) CanAll(permissions []valueobjects.Permission) bool {
	if s.IsRootOperator() || s.IsOrgRoot() {
		return true
	}
	return false
}

func (s *SessionState) CanSome(permissions []valueobjects.Permission) bool {
	if s.IsRootOperator() || s.IsOrgRoot() {
		return true
	}

	return false
}
