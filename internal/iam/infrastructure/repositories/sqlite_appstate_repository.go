package repositories

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/ports"
	"desktop/internal/shared/db"
)

type SqliteAppStateRepository struct {
	queries                *db.Queries
	OrganizationRepository ports.OrganizationRepository
	OperatorRepository     ports.OperatorRepository
}

func NewSqliteAppStateRepository(
	queries *db.Queries,
	organizationRepository ports.OrganizationRepository,
	operatorRepository ports.OperatorRepository,
) *SqliteAppStateRepository {
	return &SqliteAppStateRepository{
		queries:                queries,
		OrganizationRepository: organizationRepository,
		OperatorRepository:     operatorRepository,
	}
}

func (sql *SqliteAppStateRepository) Get(ctx context.Context) (*aggregates.AppState, error) {
	dbAppState, err := sql.queries.FindAppState(ctx)

	if err != nil {
		return nil, err
	}

	var activeOperator *aggregates.Operator = nil
	var activeOrganization *aggregates.Organization = nil

	if dbAppState.ActiveOrganizationID.Valid {
		organization, err := sql.OrganizationRepository.FindOneByID(ctx, dbAppState.ActiveOrganizationID.String)

		if err != nil {
			return nil, err
		}

		activeOrganization = organization
	}

	if dbAppState.ActiveOperatorID.Valid {
		operator, err := sql.OperatorRepository.FindOneByID(ctx, dbAppState.ActiveOperatorID.String)

		if err != nil {
			return nil, err
		}

		activeOperator = operator
	}

	currentAppState := AppStateToDomain(dbAppState, activeOperator, activeOrganization)

	return &currentAppState, nil

}

func (sql *SqliteAppStateRepository) Update(ctx context.Context, appState *aggregates.AppState) error {
	dbAppState := AppStateFromDomain(appState)

	return sql.queries.UpdateAppState(ctx, db.UpdateAppStateParams{
		ActiveOrganizationID: dbAppState.ActiveOrganizationID,
		ActiveOperatorID:     dbAppState.ActiveOperatorID,
		UpdatedAt:            dbAppState.UpdatedAt,
	})
}
