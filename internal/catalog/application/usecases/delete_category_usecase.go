package usecases

import (
	"context"
	"desktop/internal/catalog/application/commands"
	"desktop/internal/catalog/domain/ports"
)

type DeleteCategoryUseCase struct {
	categoryRepository ports.CategoryRepository
}

func NewDeleteCategoryUseCase(categoryRepository ports.CategoryRepository) *DeleteCategoryUseCase {
	return &DeleteCategoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (d *DeleteCategoryUseCase) Execute(
	ctx context.Context,
	organizationId string,
	cmd *commands.DeleteManyCategories,
) error {
	return d.categoryRepository.DeleteMany(ctx, organizationId, cmd.IDs)
}
