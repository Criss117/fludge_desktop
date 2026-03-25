package responses

import (
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/shared/db/dbutils"
)

type OperatorOrganization struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type Operator struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	Pin          string `json:"pin"`
	OperatorType string `json:"operatorType"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	DeletedAt    *int64 `json:"deletedAt"`
}

func OperatorResponseFromDomain(operator *aggregates.Operator) *Operator {
	return &Operator{
		ID:           operator.ID,
		Name:         operator.Name,
		Email:        operator.Email.Value(),
		Username:     operator.Username,
		Pin:          operator.Pin.Value(),
		OperatorType: operator.OperatorType.ToValue(),
		CreatedAt:    dbutils.TimeToInt64(operator.CreatedAt),
		UpdatedAt:    dbutils.TimeToInt64(operator.UpdatedAt),
		DeletedAt:    dbutils.TimeToInt64Nullable(operator.DeletedAt),
	}
}
