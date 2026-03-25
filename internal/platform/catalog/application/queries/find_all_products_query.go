package queries

import (
	"context"
	"desktop/internal/platform/catalog/application/responses"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

type FindAllProducts struct {
	queries *db.Queries
}

func NewFindAllProducts(queries *db.Queries) *FindAllProducts {
	return &FindAllProducts{
		queries: queries,
	}
}

func (u *FindAllProducts) Execute(ctx context.Context, organizationId string) ([]*responses.Product, error) {
	products, err := u.queries.FindAllProducts(ctx, organizationId)

	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, nil
	}

	dbProducts := make([]*responses.Product, len(products))

	for i, product := range products {
		dbProducts[i] = &responses.Product{
			ID:             product.ID,
			Sku:            product.Sku,
			Name:           product.Name,
			Description:    dbutils.StringFromSQLNullable(product.Description),
			WholesalePrice: product.WholesalePrice,
			SalePrice:      product.SalePrice,
			CostPrice:      product.CostPrice,
			OrganizationID: product.OrganizationID,
			CategoryID:     dbutils.StringFromSQLNullable(product.CategoryID),
			SupplierID:     dbutils.StringFromSQLNullable(product.SupplierID),
			CreatedAt:      product.CreatedAt,
			UpdatedAt:      product.UpdatedAt,
			DeletedAt:      dbutils.TimeToInt64Nullable(dbutils.TimeFromSQLNullable(product.DeletedAt)),
			Stock:          product.Stock,
			MinStock:       product.MinStock,
		}
	}

	return dbProducts, nil
}
