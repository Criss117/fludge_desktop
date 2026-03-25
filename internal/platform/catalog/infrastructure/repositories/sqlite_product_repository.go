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

func (r *SqliteProductRepository) FindOneById(ctx context.Context, organizationId, productId string) (*aggregates.Product, error) {
	dbProduct, err := r.queries.FindOneProduct(ctx, db.FindOneProductParams{
		ProductID:      productId,
		OrganizationID: organizationId,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return mappers.MapProductToDomain(dbProduct), nil
}

func (r *SqliteProductRepository) FindAll(ctx context.Context, organizationId string) ([]*aggregates.Product, error) {
	dbProducts, err := r.queries.FindAllProducts(ctx, organizationId)

	if err != nil {
		return nil, err
	}

	if len(dbProducts) == 0 {
		return nil, nil
	}

	products := make([]*aggregates.Product, len(dbProducts))

	for i, dbProduct := range dbProducts {
		products[i] = mappers.MapProductToDomain(dbProduct)
	}

	return products, nil
}
