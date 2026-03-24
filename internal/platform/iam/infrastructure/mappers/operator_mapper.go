package mappers

import (
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

func MapOperatorFromDomain(operator *aggregates.Operator) db.Operator {
	if operator == nil {
		return db.Operator{}
	}

	return db.Operator{
		ID:           operator.ID,
		Name:         operator.Name,
		Username:     operator.Username,
		Email:        operator.Email.Value(),
		Pin:          operator.Pin.Value(),
		OperatorType: operator.OperatorType.ToValue(),
		CreatedAt:    dbutils.TimeToInt64(operator.CreatedAt),
		UpdatedAt:    dbutils.TimeToInt64(operator.UpdatedAt),
		DeletedAt:    dbutils.TimeToSQLNullable(operator.DeletedAt),
	}
}

func MapOperatorToDomain(dbOperator db.Operator) *aggregates.Operator {
	return aggregates.ReconstituteOperator(
		dbOperator.ID,
		dbOperator.Name,
		dbOperator.Username,
		dbOperator.Email,
		dbOperator.Pin,
		dbOperator.OperatorType,
		dbutils.TimeFromInt64(dbOperator.CreatedAt),
		dbutils.TimeFromInt64(dbOperator.UpdatedAt),
		dbutils.TimeFromSQLNullable(dbOperator.DeletedAt),
	)
}
