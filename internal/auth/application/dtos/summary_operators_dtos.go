package dtos

import "desktop/internal/auth/domain"

type SummaryOperatorDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	IsRoot    bool   `json:"isRoot"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}

func DomainToSummaryOperatorDTO(operator *domain.Operator) *SummaryOperatorDTO {
	if operator == nil {
		return nil
	}

	var deletedAt *int64 = nil
	if operator.DeletedAt != nil {
		date := operator.DeletedAt.UnixMilli()
		deletedAt = &date
	}

	return &SummaryOperatorDTO{
		ID:        operator.ID,
		Name:      operator.Name,
		Username:  operator.Username,
		Email:     operator.Email,
		IsRoot:    operator.IsRoot,
		CreatedAt: operator.CreatedAt.UnixMilli(),
		UpdatedAt: operator.UpdatedAt.UnixMilli(),
		DeletedAt: deletedAt,
	}
}
