package repositories

import (
	"context"
	"desktop/internal/iam/domain/aggregates"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/platform"
)

type SqliteOperatorRepository struct {
	queries *db.Queries
}

func NewSqliteOperatorRepository(queries *db.Queries) *SqliteOperatorRepository {
	return &SqliteOperatorRepository{
		queries: queries,
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

	return OperatorToDomain(&dbOperators[0]), nil
}

func (sql *SqliteOperatorRepository) FindOneByEmail(ctx context.Context, email string) (*aggregates.Operator, error) {
	dbOperators, err := sql.queries.FindOneOperatorByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if len(dbOperators) == 0 {
		return nil, nil
	}

	return OperatorToDomain(&dbOperators[0]), nil
}

func (sql *SqliteOperatorRepository) FindOneByUsername(ctx context.Context, username string) (*aggregates.Operator, error) {
	dbOperators, err := sql.queries.FindOneOperatorByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	if len(dbOperators) == 0 {
		return nil, nil
	}

	return OperatorToDomain(&dbOperators[0]), nil
}

func (sql *SqliteOperatorRepository) Create(ctx context.Context, operator *aggregates.Operator) error {
	operatorValues := operator.ToValues()

	err := sql.queries.CreateOperator(ctx, db.CreateOperatorParams{
		ID:        operatorValues.ID,
		Name:      operatorValues.Name,
		Username:  operatorValues.Username,
		Email:     operatorValues.Email,
		Pin:       operatorValues.Pin,
		IsRoot:    platform.BoolToInt(operatorValues.IsRoot),
		CreatedAt: platform.ToMillis(operatorValues.CreatedAt),
		UpdatedAt: platform.ToMillis(operatorValues.UpdatedAt),
	})

	if err != nil {
		return err
	}

	return nil
}
