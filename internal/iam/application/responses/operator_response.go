package responses

import (
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/shared/db/platform"
)

type OperatorResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Pin       string `json:"pin"`
	IsRoot    bool   `json:"isRoot"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}

func OperatorResponseFromDomain(operator *aggregates.Operator) *OperatorResponse {
	primitiveOperator := operator.ToValues()

	return &OperatorResponse{
		ID:        primitiveOperator.ID,
		Name:      primitiveOperator.Name,
		Email:     primitiveOperator.Email,
		Username:  primitiveOperator.Username,
		Pin:       primitiveOperator.Pin,
		IsRoot:    primitiveOperator.IsRoot,
		CreatedAt: platform.TimeToInt64(primitiveOperator.CreatedAt),
		UpdatedAt: platform.TimeToInt64(primitiveOperator.UpdatedAt),
		DeletedAt: platform.TimeToInt64Nullable(primitiveOperator.DeletedAt),
	}
}
