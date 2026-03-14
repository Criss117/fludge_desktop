package infrastructure

import (
	"context"
	"desktop/internal/auth/domain"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/platform"
)

type SQLiteOperatorRepository struct {
	ctx     *context.Context
	queries *db.Queries
}

func NewSQLiteOperatorRepository(ctx *context.Context, queries *db.Queries) *SQLiteOperatorRepository {
	return &SQLiteOperatorRepository{
		ctx:     ctx,
		queries: queries,
	}
}

func (sql *SQLiteOperatorRepository) FinAll() ([]*domain.Operator, error) {
	dbAllOperators, err := sql.queries.FindAllOperators(*sql.ctx)

	if err != nil {
		return nil, err
	}

	var allOperators = make([]*domain.Operator, len(dbAllOperators))

	for i, operator := range dbAllOperators {
		allOperators[i] = OperatorToDomain(operator)
	}

	return allOperators, nil
}

func (sql *SQLiteOperatorRepository) FindOneByUsername(username string) (*domain.Operator, error) {
	dbOperators, err := sql.queries.FindOneOperatorByUsername(*sql.ctx, username)

	if err != nil {
		return nil, err
	}

	if len(dbOperators) == 0 {
		return nil, domain.ErrOperatorNotFoundByUsername()
	}

	dbOperator := dbOperators[0]

	return OperatorToDomain(dbOperator), nil
}

func (sql *SQLiteOperatorRepository) FindOneByEmail(email string) (*domain.Operator, error) {
	dbOperators, err := sql.queries.FindOneOperatorByEmail(*sql.ctx, email)

	if err != nil {
		return nil, err
	}

	if len(dbOperators) == 0 {
		return nil, domain.ErrOperatorNotFoundByEmail()
	}

	dbOperator := dbOperators[0]

	return OperatorToDomain(dbOperator), nil
}

func (sql *SQLiteOperatorRepository) FindManyByUsernameOrEmail(username, email string) ([]*domain.Operator, error) {
	dbOperators, err := sql.queries.FindManyOperatorsByEmailOrUsername(*sql.ctx, db.FindManyOperatorsByEmailOrUsernameParams{
		Email:    email,
		Username: username,
	})

	if err != nil {
		return nil, err
	}

	var allOperators = make([]*domain.Operator, len(dbOperators))

	for i, operator := range dbOperators {
		allOperators[i] = OperatorToDomain(operator)
	}

	return allOperators, nil

}

func (sql *SQLiteOperatorRepository) Create(operator domain.Operator) error {
	err := sql.queries.CreateOperator(*sql.ctx, db.CreateOperatorParams{
		ID:        operator.ID,
		Name:      operator.Name,
		Username:  operator.Username,
		Email:     operator.Email,
		Pin:       operator.Pin,
		IsRoot:    platform.BoolToInt(operator.IsRoot),
		CreatedAt: platform.ToMillis(operator.CreatedAt),
		UpdatedAt: platform.ToMillis(operator.UpdatedAt),
	})

	if err != nil {
		return err
	}

	return nil
}
