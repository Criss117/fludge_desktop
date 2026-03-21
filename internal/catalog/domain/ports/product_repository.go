package ports

import (
	"context"
	"desktop/internal/catalog/domain/aggregates"
)

type ProductRepository interface {
	Create(ctx context.Context, product *aggregates.Product) error
	Update(ctx context.Context, product *aggregates.Product) error
	// Delete(ctx context.Context, product *aggregates.Product) error

	FindOneBySku(ctx context.Context, organizationId, sku string) (*aggregates.Product, error)
	FindOneByName(ctx context.Context, organizationId, name string) (*aggregates.Product, error)
	FindOneById(ctx context.Context, organizationId, productId string) (*aggregates.Product, error)
	FindAllProducts(ctx context.Context, organizationId string) ([]*aggregates.Product, error)
}
