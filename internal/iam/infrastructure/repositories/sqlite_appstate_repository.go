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

	if dbAppState.ActiveOperatorID.Valid {
		operator, err := sql.OperatorRepository.FindOneByID(ctx, dbAppState.ActiveOperatorID.String)

		if err != nil {
			return nil, err
		}

		activeOperator = operator
	}

	currentAppState := AppStateToDomain(&dbAppState, activeOperator)

	return &currentAppState, nil

}
