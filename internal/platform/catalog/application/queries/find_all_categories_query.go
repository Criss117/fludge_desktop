package queries

import (
	"context"
	"desktop/internal/platform/catalog/application/responses"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

type FindAllCategories struct {
	queries *db.Queries
}

func NewFindAllCategories(queries *db.Queries) *FindAllCategories {
	return &FindAllCategories{
		queries: queries,
	}
}

func (u *FindAllCategories) Execute(ctx context.Context, organizationId string) ([]*responses.Category, error) {
	categories, err := u.queries.FindAllCategories(ctx, organizationId)

	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, nil
	}

	dbCategories := make([]*responses.Category, len(categories))

	for i, category := range categories {
		dbCategories[i] = &responses.Category{
			ID:             category.ID,
			Name:           category.Name,
			Description:    dbutils.StringFromSQLNullable(category.Description),
			OrganizationID: category.OrganizationID,
			CreatedAt:      category.CreatedAt,
			UpdatedAt:      category.UpdatedAt,
			DeletedAt:      dbutils.TimeToInt64Nullable(dbutils.TimeFromSQLNullable(category.DeletedAt)),
		}
	}

	return dbCategories, nil
}
