package ports

import (
	"context"
	"desktop/internal/platform/iam/domain/aggregates"
)

type OperatorRepository interface {
	Create(ctx context.Context, operator *aggregates.Operator) error
	Update(ctx context.Context, operator *aggregates.Operator) error
	Delete(ctx context.Context, operator *aggregates.Operator) error
	HardDelete(ctx context.Context, operator *aggregates.Operator) error

	FindOneByEmail(ctx context.Context, email string) (*aggregates.Operator, error)
	FindOneByUsername(ctx context.Context, username string) (*aggregates.Operator, error)
	FindOneByID(ctx context.Context, operatorId string) (*aggregates.Operator, error)
}
