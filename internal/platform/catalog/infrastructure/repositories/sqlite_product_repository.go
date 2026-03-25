package repositories

import (
	"context"
	"desktop/internal/platform/catalog/domain/aggregates"
	"desktop/internal/platform/catalog/domain/ports"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

type SqliteProductRepository struct {
	queries *db.Queries
}

func NewSqliteProductRepository(queries *db.Queries) *SqliteProductRepository {
	return &SqliteProductRepository{
		queries: queries,
	}
}

func (r *SqliteProductRepository) Create(ctx context.Context, product *aggregates.Product) error {
	if errDb := r.queries.CreateProduct(ctx, db.CreateProductParams{
		ID:             product.ID,
		Sku:            product.Sku.Value(),
		Name:           product.Name,
		Description:    dbutils.StringToSQLNullable(product.Description),
		WholesalePrice: product.WholesalePrice.Amount(),
		SalePrice:      product.SalePrice.Amount(),
		CostPrice:      product.CostPrice.Amount(),
		CategoryID:     dbutils.StringToSQLNullable(product.CategoryID),
		SupplierID:     dbutils.StringToSQLNullable(product.SupplierID),
		OrganizationID: product.OrganizationID,
		CreatedAt:      dbutils.TimeToInt64(product.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(product.UpdatedAt),
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteProductRepository) Update(ctx context.Context, product *aggregates.Product) error {
	if errDb := r.queries.UpdateProduct(ctx, db.UpdateProductParams{
		ID:             product.ID,
		Sku:            product.Sku.Value(),
		Name:           product.Name,
		Description:    dbutils.StringToSQLNullable(product.Description),
		WholesalePrice: product.WholesalePrice.Amount(),
		SalePrice:      product.SalePrice.Amount(),
		CostPrice:      product.CostPrice.Amount(),
		CategoryID:     dbutils.StringToSQLNullable(product.CategoryID),
		SupplierID:     dbutils.StringToSQLNullable(product.SupplierID),
		UpdatedAt:      dbutils.TimeToInt64(product.UpdatedAt),
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteProductRepository) Delete(ctx context.Context, product *aggregates.Product) error {
	if errDb := r.queries.DeleteProduct(ctx, db.DeleteProductParams{
		ProductID:      product.ID,
		OrganizationID: product.OrganizationID,
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteProductRepository) ExistsBy(ctx context.Context, organizationId string, params ports.ExistsByParams) (ports.ExistsByReturn, error) {
	products, err := r.queries.ExistsProductByNameOrSku(ctx, db.ExistsProductByNameOrSkuParams{
		Name:           params.Name,
		Sku:            params.Sku,
		OrganizationID: organizationId,
	})

	if err != nil {
		return ports.ExistsByReturn{}, err
	}

	if len(products) == 0 {
		return ports.ExistsByReturn{
			Name: false,
			Sku:  false,
		}, nil
	}

	var existsByName bool = false
	var existsBySku bool = false

	for _, product := range products {
		if product.Name == params.Name {
			existsByName = true
		}

		if product.Sku == params.Sku {
			existsBySku = true
		}
	}

	return ports.ExistsByReturn{
		Name: existsByName,
		Sku:  existsBySku,
	}, nil
}
