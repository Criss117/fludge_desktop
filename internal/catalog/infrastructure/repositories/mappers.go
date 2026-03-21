package repositories

import (
	"desktop/internal/catalog/domain/aggregates"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/platform"
)

func ProductToDomain(dbProduct *db.Product) *aggregates.Product {
	return aggregates.ReconstituteProduct(
		dbProduct.ID,
		dbProduct.Sku,
		dbProduct.Name,
		platform.FromStringNullable(dbProduct.Description),
		dbProduct.WholesalePrice,
		dbProduct.SalePrice,
		dbProduct.CostPrice,
		dbProduct.Stock,
		dbProduct.MinStock,
		platform.FromStringNullable(dbProduct.CategoryID),
		dbProduct.OrganizationID,
		platform.FromStringNullable(dbProduct.SupplierID),
		platform.FromMillis(dbProduct.CreatedAt),
		platform.FromMillis(dbProduct.UpdatedAt),
		platform.FromMillisNullable(dbProduct.DeletedAt),
	)
}

func CategoryToDomain(dbCategory *db.Category) *aggregates.Category {
	return aggregates.ReconstituteCategory(
		dbCategory.ID,
		dbCategory.Name,
		platform.FromStringNullable(dbCategory.Description),
		dbCategory.OrganizationID,
		platform.FromMillis(dbCategory.CreatedAt),
		platform.FromMillis(dbCategory.UpdatedAt),
		platform.FromMillisNullable(dbCategory.DeletedAt),
	)
}
