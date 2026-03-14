package infrastructure

import (
	"context"
	"desktop/internal/auth/domain"
	"desktop/internal/shared/db"
)

type SQLiteAppStateRepository struct {
	ctx     *context.Context
	queries *db.Queries
}

func NewSQLiteAppStateRepository(ctx *context.Context, queries *db.Queries) *SQLiteAppStateRepository {
	return &SQLiteAppStateRepository{
		ctx:     ctx,
		queries: queries,
	}
}

func (sql *SQLiteAppStateRepository) FindAppState() (*domain.AppState, error) {
	dbAppState, err := sql.queries.FindAppState(*sql.ctx)

	if err != nil {
		return nil, err
	}

	dbOperators, errOpertators := sql.queries.FindAllOperators(*sql.ctx)

	if errOpertators != nil {
		return nil, errOpertators
	}

	domainAppState := AppStateToDomain(&dbAppState, dbOperators)

	return domainAppState, nil
}

func (sql *SQLiteAppStateRepository) Update(appState domain.AppState) error {
	dbAppState := AppStateFromDomain(&appState)

	errUpdating := sql.queries.UpdateAppState(*sql.ctx, db.UpdateAppStateParams{
		ActiveOrganizationID: dbAppState.ActiveOrganizationID,
		ActiveOperatorID:     dbAppState.ActiveOperatorID,
		UpdatedAt:            dbAppState.UpdatedAt,
	})

	return errUpdating
}
