package ports

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
)

type TeamRepository interface {
	// FindAllByOrganizationId(ctx context.Context, organizationId string) ([]*aggregates.Team, error)
	FindAllByOperatorId(ctx context.Context, operatorId string) ([]*aggregates.Team, error)
}
