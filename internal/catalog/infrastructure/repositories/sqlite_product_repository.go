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

func (r *SQLiteProductRepository) FindOneBySku(ctx context.Context, organizationId, sku string) (*aggregates.Product, error) {
	dbProducts, err := r.queries.FindOneProductBySku(ctx, db.FindOneProductBySkuParams{
		Sku:            sku,
		OrganizationID: organizationId,
	})

	if err != nil {
		return nil, err
	}

	if len(dbProducts) == 0 {
		return nil, nil
	}

	product := dbProducts[0]

	return ProductToDomain(&product), nil
}

func (r *SQLiteProductRepository) FindOneByName(ctx context.Context, organizationId, name string) (*aggregates.Product, error) {
	dbProducts, err := r.queries.FindOneProductByName(ctx, db.FindOneProductByNameParams{
		OrganizationID: organizationId,
		LOWER:          name,
	})

	if err != nil {
		return nil, err
	}

	if len(dbProducts) == 0 {
		return nil, nil
	}

	product := dbProducts[0]

	return ProductToDomain(&product), nil
}

func (r *SQLiteProductRepository) FindAllProducts(ctx context.Context, organizationId string) ([]*aggregates.Product, error) {
	dbProducts, err := r.queries.FindAllProductsByOrganizationId(ctx, organizationId)

	if err != nil {
		return nil, err
	}

	if len(dbProducts) == 0 {
		return nil, derrors.ErrProductNotFound
	}

	products := make([]*aggregates.Product, len(dbProducts))

	for index, dbProduct := range dbProducts {
		products[index] = ProductToDomain(&dbProduct)
	}

	return products, nil
}

func (r *SQLiteProductRepository) FindOneById(
	ctx context.Context,
	organizationId, productId string,
) (*aggregates.Product, error) {
	dbProducts, err := r.queries.FindOneProductById(ctx, db.FindOneProductByIdParams{
		ID:             productId,
		OrganizationID: organizationId,
	})

	if err != nil {
		return nil, err
	}

	if len(dbProducts) == 0 {
		return nil, nil
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

func (r *SQLiteProductRepository) Update(ctx context.Context, product *aggregates.Product) error {
	return r.queries.UpdateProduct(ctx, db.UpdateProductParams{
		Sku:            product.Sku.Value(),
		Name:           product.Name,
		Description:    platform.ToStringNullable(product.Description),
		WholesalePrice: product.WholesalePrice.Amount(),
		SalePrice:      product.SalePrice.Amount(),
		CostPrice:      product.CostPrice.Amount(),
		Stock:          product.Stock,
		MinStock:       product.MinStock,
		CategoryID:     platform.ToStringNullable(product.CategoryID),
		SupplierID:     platform.ToStringNullable(product.SupplierID),
		UpdatedAt:      platform.ToMillis(product.UpdatedAt),
		ID:             product.ID,
		OrganizationID: product.OrganizationID,
	})
}
