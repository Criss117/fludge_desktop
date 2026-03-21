package ports

import (
	"context"
	"desktop/internal/catalog/domain/aggregates"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, organizationId string) ([]*aggregates.Category, error)
	FindOneById(ctx context.Context, organizationId string, categoryId string) (*aggregates.Category, error)
	FindOneByName(ctx context.Context, organizationId string, name string) (*aggregates.Category, error)

	Create(ctx context.Context, organizationId string, category *aggregates.Category) error
	Update(ctx context.Context, organizationId string, category *aggregates.Category) error
	Delete(ctx context.Context, organizationId string, categoryId string) error
	DeleteMany(ctx context.Context, organizationId string, categoryIds []string) error
}
