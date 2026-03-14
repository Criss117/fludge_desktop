package infrastructure

import (
	"desktop/internal/auth/domain"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/platform"
)

func OperatorToDomain(operator db.Operator) *domain.Operator {
	return &domain.Operator{
		ID:        operator.ID,
		Name:      operator.Name,
		Username:  operator.Username,
		Pin:       operator.Pin,
		Email:     operator.Email,
		IsRoot:    platform.IntToBool(operator.IsRoot),
		CreatedAt: platform.FromMillis(operator.CreatedAt),
		UpdatedAt: platform.FromMillis(operator.UpdatedAt),
		DeletedAt: platform.FromMillisNullable(operator.DeletedAt),
	}
}

func OperatorFromDomain(operator domain.Operator) *db.Operator {
	return &db.Operator{
		ID:        operator.ID,
		Name:      operator.Name,
		Username:  operator.Username,
		Pin:       operator.Pin,
		Email:     operator.Email,
		IsRoot:    platform.BoolToInt(operator.IsRoot),
		CreatedAt: platform.ToMillis(operator.CreatedAt),
		UpdatedAt: platform.ToMillis(operator.UpdatedAt),
		DeletedAt: platform.ToMillisNullable(operator.DeletedAt),
	}
}

func AppStateToDomain(appState *db.AppState, dbOperators []db.Operator) *domain.AppState {
	domainOperators := make([]*domain.Operator, len(dbOperators))

	var activeOperator *domain.Operator = nil

	for i, operator := range dbOperators {
		domainOperators[i] = OperatorToDomain(operator)
	}

	if appState.ActiveOperatorID.Valid {
		for _, operator := range domainOperators {
			if operator.ID == appState.ActiveOperatorID.String {
				activeOperator = operator
				break
			}
		}
	}

	return &domain.AppState{
		ActiveOrganizationID: platform.FromStringNullable(appState.ActiveOrganizationID),
		ActiveOperator:       activeOperator,
		UpdatedAt:            platform.FromMillis(appState.UpdatedAt),
		Operators:            domainOperators,
	}
}

func AppStateFromDomain(appState *domain.AppState) *db.AppState {
	return &db.AppState{
		ActiveOrganizationID: platform.ToStringNullable(appState.ActiveOrganizationID),
		ActiveOperatorID:     platform.ToStringNullable(&appState.ActiveOperator.ID),
		UpdatedAt:            platform.ToMillis(appState.UpdatedAt),
	}
}
