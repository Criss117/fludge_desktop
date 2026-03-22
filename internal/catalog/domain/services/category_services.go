package services

import (
	"context"
	"desktop/internal/catalog/domain/ports"
)

type CategoryService struct {
	categoryRepository ports.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) EnsureCategoryExists(ctx context.Context, organizationId, categoryId string) (bool, error) {
	category, err := s.categoryRepository.FindOneById(ctx, organizationId, categoryId)

	if err != nil {
		return false, err
	}

	if category != nil {
		return true, nil
	}

	return false, nil
}
