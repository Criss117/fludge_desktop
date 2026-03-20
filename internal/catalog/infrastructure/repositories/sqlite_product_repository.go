package repositories

import (
	"context"
	"desktop/internal/catalog/domain/aggregates"
	"desktop/internal/catalog/domain/derrors"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/platform"
)

type SQLiteProductRepository struct {
	queries *db.Queries
}

func NewSQLiteProductRepository(queries *db.Queries) *SQLiteProductRepository {
	return &SQLiteProductRepository{
		queries: queries,
	}
}

func (r *SQLiteProductRepository) FindOneBySku(ctx context.Context, sku string) (*aggregates.Product, error) {
	dbProducts, err := r.queries.FindOneProductBySku(ctx, sku)

	if err != nil {
		return nil, err
	}

	if len(dbProducts) == 0 {
		return nil, derrors.ErrProductNotFound
	}

	product := dbProducts[0]

	return ProductToDomain(&product), nil
}

func (r *SQLiteProductRepository) FindOneByName(ctx context.Context, name string) (*aggregates.Product, error) {
	dbProducts, err := r.queries.FindOneProductByName(ctx, name)

	if err != nil {
		return nil, err
	}

	if len(dbProducts) == 0 {
		return nil, derrors.ErrProductNotFound
	}

	product := dbProducts[0]

	return ProductToDomain(&product), nil
}

func (r *SQLiteProductRepository) Create(ctx context.Context, product *aggregates.Product) error {
	return r.queries.CreateProduct(ctx, db.CreateProductParams{
		ID:             product.ID,
		Sku:            product.Sku.Value(),
		Name:           product.Name,
		Description:    platform.ToStringNullable(product.Description),
		WholesalePrice: product.WholesalePrice.Amount(),
		SalePrice:      product.SalePrice.Amount(),
		CostPrice:      product.CostPrice.Amount(),
		Stock:          product.Stock,
		MinStock:       product.MinStock,
		CategoryID:     platform.ToStringNullable(product.CategoryID),
		OrganizationID: product.OrganizationID,
		SupplierID:     platform.ToStringNullable(product.SupplierID),
		CreatedAt:      platform.ToMillis(product.CreatedAt),
		UpdatedAt:      platform.ToMillis(product.UpdatedAt),
	})
}
