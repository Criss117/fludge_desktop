package repositories

import (
	"context"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/infrastructure/mappers"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

type SqliteAppRepository struct {
	queries *db.Queries
}

func NewSqliteAppRepository(queries *db.Queries) *SqliteAppRepository {
	return &SqliteAppRepository{
		queries: queries,
	}
}

func (r *SqliteAppRepository) Get(ctx context.Context) (*aggregates.AppState, error) {
	dbAppState, errDb := r.queries.FindAppState(ctx)

	if errDb != nil {
		return nil, errDb
	}

	return mappers.MapAppStateToDomain(dbAppState), nil
}

func (r *SqliteAppRepository) Update(ctx context.Context, appState *aggregates.AppState) error {
	if errDb := r.queries.UpdateAppState(ctx, db.UpdateAppStateParams{
		ActiveOrganizationID: dbutils.StringToSQLNullable(appState.ActiveOrganizationID),
		ActiveOperatorID:     dbutils.StringToSQLNullable(appState.ActiveOperatorID),
		UpdatedAt:            dbutils.TimeToInt64(appState.UpdatedAt),
	}); errDb != nil {
		return errDb
	}
	return nil
}
