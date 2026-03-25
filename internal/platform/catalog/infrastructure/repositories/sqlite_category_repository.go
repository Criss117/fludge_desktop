package repositories

import (
	"context"
	"database/sql"
	"desktop/internal/platform/catalog/domain/aggregates"
	"desktop/internal/platform/catalog/infrastructure/mappers"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
	"errors"
)

type SqliteCategoryRepository struct {
	queries *db.Queries
}

func NewSqliteCategoryRepository(queries *db.Queries) *SqliteCategoryRepository {
	return &SqliteCategoryRepository{
		queries: queries,
	}
}

func (r *SqliteCategoryRepository) Create(ctx context.Context, category *aggregates.Category) error {
	if errDb := r.queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:             category.ID,
		Name:           category.Name,
		Description:    dbutils.StringToSQLNullable(category.Description),
		OrganizationID: category.OrganizationID,
		CreatedAt:      dbutils.TimeToInt64(category.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(category.UpdatedAt),
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteCategoryRepository) Update(ctx context.Context, category *aggregates.Category) error {
	if errDb := r.queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		CategoryID:     category.ID,
		Name:           category.Name,
		Description:    dbutils.StringToSQLNullable(category.Description),
		OrganizationID: category.OrganizationID,
		UpdatedAt:      dbutils.TimeToInt64(category.UpdatedAt),
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteCategoryRepository) Delete(ctx context.Context, category *aggregates.Category) error {
	if errDb := r.queries.DeleteCategory(ctx, db.DeleteCategoryParams{
		CategoryID:     category.ID,
		OrganizationID: category.OrganizationID,
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteCategoryRepository) DeleteMany(ctx context.Context, organizationId string, categoryIds []string) error {
	if errDb := r.queries.DeleteCategories(ctx, db.DeleteCategoriesParams{
		CategoryIds:    categoryIds,
		OrganizationID: organizationId,
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteCategoryRepository) FindOneById(ctx context.Context, organizationId, categoryId string) (*aggregates.Category, error) {
	category, err := r.queries.FindOneCategory(ctx, db.FindOneCategoryParams{
		CategoryID:     categoryId,
		OrganizationID: organizationId,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return mappers.MapCategoryToDomain(category), nil
}

func (r *SqliteCategoryRepository) FindAll(ctx context.Context, organizationId string) ([]*aggregates.Category, error) {
	dbCategories, err := r.queries.FindAllCategories(ctx, organizationId)

	if err != nil {
		return nil, err
	}

	if len(dbCategories) == 0 {
		return nil, nil
	}

	categories := make([]*aggregates.Category, len(dbCategories))

	for i, dbCategory := range dbCategories {
		categories[i] = mappers.MapCategoryToDomain(dbCategory)
	}

	return categories, nil
}

func (r *SqliteCategoryRepository) FindOneByName(ctx context.Context, organizationId, name string) (*aggregates.Category, error) {
	category, err := r.queries.FindOneCategoryByName(ctx, db.FindOneCategoryByNameParams{
		CategoryName:   name,
		OrganizationID: organizationId,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return mappers.MapCategoryToDomain(category), nil
}
