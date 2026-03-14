package dtos

import "desktop/internal/auth/domain"

type AppStateDTO struct {
	ActiveOrganizationID *string               `json:"activeOrganizationId"`
	ActiveOperator       *SummaryOperatorDTO   `json:"activeOperator"`
	UpdatedAt            int64                 `json:"updatedAt"`
	Operators            []*SummaryOperatorDTO `json:"operators"`
}

func DomainToAppStateDTO(appState *domain.AppState) *AppStateDTO {
	var operators = make([]*SummaryOperatorDTO, len(appState.Operators))
	for i, operator := range appState.Operators {
		operators[i] = DomainToSummaryOperatorDTO(operator)
	}

	var activeOperator *SummaryOperatorDTO
	if appState.ActiveOperator != nil {
		activeOperator = DomainToSummaryOperatorDTO(appState.ActiveOperator)
	}

	return &AppStateDTO{
		ActiveOrganizationID: appState.ActiveOrganizationID,
		ActiveOperator:       activeOperator,
		UpdatedAt:            appState.UpdatedAt.UnixMilli(),
		Operators:            operators,
	}
}
