package mappers

import (
	"desktop/internal/platform/catalog/domain/aggregates"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

func MapCategoryToDomain(dbCategory db.Category) *aggregates.Category {
	return aggregates.ReconstituteCategory(
		dbCategory.ID,
		dbCategory.Name,
		dbutils.StringFromSQLNullable(dbCategory.Description),
		dbCategory.OrganizationID,
		dbutils.TimeFromInt64(dbCategory.CreatedAt),
		dbutils.TimeFromInt64(dbCategory.UpdatedAt),
		dbutils.TimeFromSQLNullable(dbCategory.DeletedAt),
	)
}

func MapCategoryFromDomain(category *aggregates.Category) db.Category {
	return db.Category{
		ID:             category.ID,
		Name:           category.Name,
		Description:    dbutils.StringToSQLNullable(category.Description),
		OrganizationID: category.OrganizationID,
		CreatedAt:      dbutils.TimeToInt64(category.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(category.UpdatedAt),
		DeletedAt:      dbutils.TimeToSQLNullable(category.DeletedAt),
	}
}
