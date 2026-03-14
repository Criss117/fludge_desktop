package domain

import "time"

type AppState struct {
	ActiveOrganizationID *string
	ActiveOperator       *Operator
	UpdatedAt            time.Time
	Operators            []*Operator
}

func NewAppState(activeOrganizationID *string, operator *Operator, updatedAt time.Time, operators []*Operator) AppState {
	return AppState{
		ActiveOrganizationID: activeOrganizationID,
		ActiveOperator:       operator,
		UpdatedAt:            updatedAt,
	}
}

func (as *AppState) SetOrganization(activeOrganizationID string) {
	as.ActiveOrganizationID = &activeOrganizationID
	as.UpdatedAt = time.Now()
}

func (as *AppState) SetOperator(operator Operator) error {
	existsOperator := false

	for _, op := range as.Operators {
		if op.ID == operator.ID {
			existsOperator = true
			break
		}
	}

	if !existsOperator {
		return ErrAppStateNotSetOperator()
	}

	as.ActiveOperator = &operator
	as.UpdatedAt = time.Now()

	return nil
}
