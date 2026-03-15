package ports

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
)

type OperatorRepository interface {
	Create(ctx context.Context, operator *aggregates.Operator) error
	// Update(operator *aggregates.Operator) error
	// Delete(operator *aggregates.Operator) error

	FindOneByID(ctx context.Context, operatorId string) (*aggregates.Operator, error)
	FindOneByEmail(ctx context.Context, email string) (*aggregates.Operator, error)
	FindOneByUsername(ctx context.Context, username string) (*aggregates.Operator, error)
	// ExistsByEmail(email string) (bool, error)
	// ExistsByUsername(username string) (bool, error)
}
