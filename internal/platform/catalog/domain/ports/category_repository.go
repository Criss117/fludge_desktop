package ports

import (
	"context"
	"desktop/internal/platform/catalog/domain/aggregates"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *aggregates.Category) error
	Update(ctx context.Context, category *aggregates.Category) error
	Delete(ctx context.Context, category *aggregates.Category) error
	DeleteMany(ctx context.Context, organizationId string, categoryIds []string) error

	FindOneById(ctx context.Context, organizationId, categoryId string) (*aggregates.Category, error)
	FindOneByName(ctx context.Context, organizationId, name string) (*aggregates.Category, error)
	FindAll(ctx context.Context, organizationId string) ([]*aggregates.Category, error)
}
