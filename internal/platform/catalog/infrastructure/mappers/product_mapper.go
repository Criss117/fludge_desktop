package mappers

import (
	"desktop/internal/platform/catalog/domain/aggregates"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

func MapProductToDomain(product db.Product) *aggregates.Product {
	return aggregates.ReconstituteProduct(
		product.ID,
		product.Sku,
		product.Name,
		product.OrganizationID,
		dbutils.StringFromSQLNullable(product.Description),
		product.WholesalePrice,
		product.SalePrice,
		product.CostPrice,
		dbutils.StringFromSQLNullable(product.CategoryID),
		dbutils.StringFromSQLNullable(product.SupplierID),
		dbutils.TimeFromInt64(product.CreatedAt),
		dbutils.TimeFromInt64(product.UpdatedAt),
		dbutils.TimeFromSQLNullable(product.DeletedAt),
	)
}
