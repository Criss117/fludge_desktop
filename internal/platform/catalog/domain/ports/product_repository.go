package ports

import (
	"context"
	"desktop/internal/platform/catalog/domain/aggregates"
)

type ProductRepository interface {
	Create(ctx context.Context, product *aggregates.Product) error
	Update(ctx context.Context, product *aggregates.Product) error
	Delete(ctx context.Context, product *aggregates.Product) error

	FindOneById(ctx context.Context, organizationId, productId string) (*aggregates.Product, error)
	FindAll(ctx context.Context, organizationId string) ([]*aggregates.Product, error)
}
