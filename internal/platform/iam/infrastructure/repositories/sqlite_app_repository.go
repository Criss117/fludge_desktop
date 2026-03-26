package repositories

import (
	"context"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/domain/ports"
	"desktop/internal/platform/iam/infrastructure/mappers"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

type SqliteAppRepository struct {
	queries                *db.Queries
	organizationRepository ports.OrganizationRepository
	operatorRepository     ports.OperatorRepository
}

func NewSqliteAppRepository(
	queries *db.Queries,
	organizationRepository ports.OrganizationRepository,
	operatorRepository ports.OperatorRepository,
) *SqliteAppRepository {
	return &SqliteAppRepository{
		queries:                queries,
		organizationRepository: organizationRepository,
		operatorRepository:     operatorRepository,
	}
}

func (r *SqliteAppRepository) Get(ctx context.Context) (*aggregates.AppState, error) {
	dbAppState, errDb := r.queries.FindAppState(ctx)

	if errDb != nil {
		return nil, errDb
	}

	return mappers.MapAppStateToDomain(dbAppState), nil
}

func (r *SqliteAppRepository) GetWithAggregates(ctx context.Context) (*ports.AppStateWithAggregates, error) {
	dbAppState, errDb := r.queries.FindAppState(ctx)

	if errDb != nil {
		return nil, errDb
	}

	appState := mappers.MapAppStateToDomain(dbAppState)

	var activeOperator *aggregates.Operator = nil
	var activeOrganization *aggregates.Organization = nil

	if appState.ActiveOperatorID != nil {
		operator, err := r.operatorRepository.FindOneByID(ctx, *appState.ActiveOperatorID)

		if err != nil {
			return nil, err
		}

		activeOperator = operator
	}

	if appState.ActiveOrganizationID != nil {
		organization, err := r.organizationRepository.FindOneById(ctx, *appState.ActiveOrganizationID)

		if err != nil {
			return nil, err
		}

		activeOrganization = organization
	}

	return &ports.AppStateWithAggregates{
		AppState:     appState,
		Operator:     activeOperator,
		Organization: activeOrganization,
	}, nil
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
