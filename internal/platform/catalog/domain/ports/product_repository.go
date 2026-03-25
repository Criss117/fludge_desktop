package ports

import (
	"context"
	"desktop/internal/platform/catalog/domain/aggregates"
)

type ExistsByParams struct {
	Name      string
	Sku       string
	ProductID *string
}

type ExistsByReturn struct {
	Name bool
	Sku  bool
}

type ProductRepository interface {
	Create(ctx context.Context, product *aggregates.Product) error
	Update(ctx context.Context, product *aggregates.Product) error
	Delete(ctx context.Context, product *aggregates.Product) error

	ExistsBy(ctx context.Context, organizationId string, params ExistsByParams) (ExistsByReturn, error)
	FindOneById(ctx context.Context, organizationId, productId string) (*aggregates.Product, error)
}
