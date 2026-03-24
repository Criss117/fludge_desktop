package repositories

import (
	"context"
	"database/sql"
	"desktop/internal/platform/iam/domain/aggregates"
	"desktop/internal/platform/iam/infrastructure/mappers"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
	"desktop/internal/shared/derrors"
	"errors"
)

type SqliteOperatorRepository struct {
	queries *db.Queries
}

func NewSqliteOperatorRepository(queries *db.Queries) *SqliteOperatorRepository {
	return &SqliteOperatorRepository{
		queries: queries,
	}
}

func (r *SqliteOperatorRepository) Create(ctx context.Context, operator *aggregates.Operator) error {
	if errDb := r.queries.CreateOperator(ctx, db.CreateOperatorParams{
		ID:           operator.ID,
		Name:         operator.Name,
		Username:     operator.Username,
		Email:        operator.Email.Value(),
		Pin:          operator.Pin.Value(),
		OperatorType: operator.OperatorType.ToValue(),
		CreatedAt:    dbutils.TimeToInt64(operator.CreatedAt),
		UpdatedAt:    dbutils.TimeToInt64(operator.UpdatedAt),
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteOperatorRepository) Update(ctx context.Context, operator *aggregates.Operator) error {
	if errDb := r.queries.UpdateOperator(ctx, db.UpdateOperatorParams{
		ID:           operator.ID,
		Name:         operator.Name,
		Username:     operator.Username,
		Email:        operator.Email.Value(),
		Pin:          operator.Pin.Value(),
		OperatorType: operator.OperatorType.ToValue(),
		UpdatedAt:    dbutils.TimeToInt64(operator.UpdatedAt),
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteOperatorRepository) HardDelete(ctx context.Context, operator *aggregates.Operator) error {
	if errDb := r.queries.DeleteOperator(ctx, operator.ID); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteOperatorRepository) Delete(ctx context.Context, operator *aggregates.Operator) error {
	if operator.DeletedAt == nil {
		return derrors.ErrInvalidDeleteState
	}

	if errDb := r.queries.SoftDeleteOperator(ctx, db.SoftDeleteOperatorParams{
		DeletedAt: dbutils.TimeToSQLNullable(operator.DeletedAt),
		ID:        operator.ID,
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteOperatorRepository) FindOneByEmail(ctx context.Context, email string) (*aggregates.Operator, error) {
	dbOperator, errDb := r.queries.FindOneOperatorByEmail(ctx, email)

	if errDb != nil {
		if errors.Is(errDb, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errDb
	}

	return mappers.MapOperatorToDomain(dbOperator), nil
}

func (r *SqliteOperatorRepository) FindOneByUsername(ctx context.Context, username string) (*aggregates.Operator, error) {
	dbOperator, errDb := r.queries.FindOneOperatorByUsername(ctx, username)

	if errDb != nil {
		if errors.Is(errDb, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errDb
	}

	return mappers.MapOperatorToDomain(dbOperator), nil
}

func (r *SqliteOperatorRepository) FindOneByID(ctx context.Context, operatorId string) (*aggregates.Operator, error) {
	dbOperator, errDb := r.queries.FindOneOperatorById(ctx, operatorId)

	if errDb != nil {
		if errors.Is(errDb, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errDb
	}

	return mappers.MapOperatorToDomain(dbOperator), nil
}
