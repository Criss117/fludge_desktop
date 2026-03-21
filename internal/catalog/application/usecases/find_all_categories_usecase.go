package usecases

import (
	"context"
	"desktop/internal/catalog/domain/aggregates"
	"desktop/internal/catalog/domain/ports"
)

type FindAllCategoriesUseCase struct {
	categoriesRepository ports.CategoryRepository
}

func NewFindAllCategoriesUseCase(categoriesRepository ports.CategoryRepository) *FindAllCategoriesUseCase {
	return &FindAllCategoriesUseCase{
		categoriesRepository: categoriesRepository,
	}
}

func (u *FindAllCategoriesUseCase) Execute(ctx context.Context, organizationId string) ([]*aggregates.Category, error) {
	return u.categoriesRepository.FindAll(ctx, organizationId)
}
