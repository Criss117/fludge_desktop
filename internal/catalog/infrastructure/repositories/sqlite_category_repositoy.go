package repositories

import (
	"context"
	"desktop/internal/catalog/domain/aggregates"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/platform"
)

type SQLiteCategoryRepository struct {
	queries *db.Queries
}

func NewSQLiteCategoryRepository(queries *db.Queries) *SQLiteCategoryRepository {
	return &SQLiteCategoryRepository{
		queries: queries,
	}
}

func (r *SQLiteCategoryRepository) FindAll(ctx context.Context, organizationId string) ([]*aggregates.Category, error) {
	dbCategories, err := r.queries.FindAllCategories(ctx, organizationId)

	if err != nil {
		return nil, err
	}

	if len(dbCategories) == 0 {
		return nil, nil
	}

	categories := make([]*aggregates.Category, len(dbCategories))

	for index, dbCategory := range dbCategories {
		categories[index] = CategoryToDomain(&dbCategory)
	}

	return categories, nil
}

func (r *SQLiteCategoryRepository) FindOneById(
	ctx context.Context,
	organizationId string,
	categoryId string,
) (*aggregates.Category, error) {
	dbCategory, err := r.queries.FindOneCategoryById(ctx, db.FindOneCategoryByIdParams{
		ID:             categoryId,
		OrganizationID: organizationId,
	})

	if err != nil {
		return nil, err
	}

	if len(dbCategory) == 0 {
		return nil, nil
	}

	category := dbCategory[0]

	return CategoryToDomain(&category), nil
}

func (r *SQLiteCategoryRepository) FindOneByName(
	ctx context.Context,
	organizationId string,
	name string,
) (*aggregates.Category, error) {
	dbCategory, err := r.queries.FindOneCategoryByName(ctx, db.FindOneCategoryByNameParams{
		LOWER:          name,
		OrganizationID: organizationId,
	})

	if err != nil {
		return nil, err
	}

	if len(dbCategory) == 0 {
		return nil, nil
	}

	category := dbCategory[0]

	return CategoryToDomain(&category), nil
}

func (r *SQLiteCategoryRepository) Create(ctx context.Context, organizationId string, category *aggregates.Category) error {
	return r.queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:             category.ID,
		Name:           category.Name,
		Description:    platform.ToStringNullable(category.Description),
		OrganizationID: organizationId,
		CreatedAt:      platform.ToMillis(category.CreatedAt),
		UpdatedAt:      platform.ToMillis(category.UpdatedAt),
	})
}

func (r *SQLiteCategoryRepository) Update(ctx context.Context, organizationId string, category *aggregates.Category) error {
	return r.queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		Name:           category.Name,
		Description:    platform.ToStringNullable(category.Description),
		UpdatedAt:      platform.ToMillis(category.UpdatedAt),
		ID:             category.ID,
		OrganizationID: organizationId,
	})
}

func (r *SQLiteCategoryRepository) Delete(ctx context.Context, organizationId string, categoryId string) error {
	return r.queries.DeleteCategory(ctx, db.DeleteCategoryParams{
		ID:             categoryId,
		OrganizationID: organizationId,
	})
}
