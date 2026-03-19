package repositories

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/iam/domain/ports"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/platform"
)

type SqliteOperatorRepository struct {
	queries           *db.Queries
	organizationRepos ports.OrganizationRepository
}

func NewSqliteOperatorRepository(queries *db.Queries, organizationRepos ports.OrganizationRepository) *SqliteOperatorRepository {
	return &SqliteOperatorRepository{
		queries:           queries,
		organizationRepos: organizationRepos,
	}
}

func (sql *SqliteOperatorRepository) FindOneByID(ctx context.Context, operatorId string) (*aggregates.Operator, error) {
	dbOperators, err := sql.queries.FindOneOperatorById(ctx, operatorId)

	if err != nil {
		return nil, err
	}

	if len(dbOperators) == 0 {
		return nil, nil
	}

	operatorOrganizations, err := sql.organizationRepos.FindByOperator(ctx, operatorId)

	return OperatorToDomain(&dbOperators[0], operatorOrganizations), nil
}

func (sql *SqliteOperatorRepository) FindOneByEmail(ctx context.Context, email string) (*aggregates.Operator, error) {
	dbOperators, err := sql.queries.FindOneOperatorByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if len(dbOperators) == 0 {
		return nil, nil
	}

	operator := dbOperators[0]

	operatorOrganizations, err := sql.organizationRepos.FindByOperator(ctx, operator.ID)

	return OperatorToDomain(&operator, operatorOrganizations), nil
}

func (sql *SqliteOperatorRepository) FindOneByUsername(ctx context.Context, username string) (*aggregates.Operator, error) {
	dbOperators, err := sql.queries.FindOneOperatorByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	if len(dbOperators) == 0 {
		return nil, nil
	}

	operator := dbOperators[0]

	operatorOrganizations, err := sql.organizationRepos.FindByOperator(ctx, operator.ID)

	return OperatorToDomain(&operator, operatorOrganizations), nil
}

func (sql *SqliteOperatorRepository) Create(ctx context.Context, operator *aggregates.Operator) error {
	err := sql.queries.CreateOperator(ctx, db.CreateOperatorParams{
		ID:        operator.ID,
		Name:      operator.Name,
		Username:  operator.Username,
		Email:     operator.Email.Value(),
		Pin:       operator.Pin.Value(),
		IsRoot:    platform.BoolToInt(operator.Root),
		CreatedAt: platform.ToMillis(operator.CreatedAt),
		UpdatedAt: platform.ToMillis(operator.UpdatedAt),
	})

	if err != nil {
		return err
	}

	return nil
}
