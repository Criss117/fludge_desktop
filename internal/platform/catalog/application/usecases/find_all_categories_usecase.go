package usecases

import (
	"context"
	"desktop/internal/platform/catalog/domain/aggregates"
	"desktop/internal/platform/catalog/domain/ports"
)

type FindAllCategories struct {
	categoryrepository ports.CategoryRepository
}

func NewFindAllCategories(categoryrepository ports.CategoryRepository) *FindAllCategories {
	return &FindAllCategories{
		categoryrepository: categoryrepository,
	}
}

func (u *FindAllCategories) Execute(ctx context.Context, organizationId string) ([]*aggregates.Category, error) {
	categories, err := u.categoryrepository.FindAll(ctx, organizationId)

	if err != nil {
		return nil, err
	}

	return categories, nil
}
