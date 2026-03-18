package ports

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
)

type AppStateRepository interface {
	Get(ctx context.Context) (*aggregates.AppState, error)
	Update(ctx context.Context, appState *aggregates.AppState) error
}
