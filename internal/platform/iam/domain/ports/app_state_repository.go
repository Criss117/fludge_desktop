package ports

import (
	"context"
	"desktop/internal/platform/iam/domain/aggregates"
)

type AppStateWithAggregates struct {
	AppState     *aggregates.AppState
	Operator     *aggregates.Operator
	Organization *aggregates.Organization
}

type AppStateRepository interface {
	Get(ctx context.Context) (*aggregates.AppState, error)
	GetWithAggregates(ctx context.Context) (*AppStateWithAggregates, error)
	Update(ctx context.Context, appState *aggregates.AppState) error
}
