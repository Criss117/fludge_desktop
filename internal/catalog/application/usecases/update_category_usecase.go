package usecases

import (
	"context"
	"desktop/internal/catalog/application/commands"
	"desktop/internal/catalog/domain/aggregates"
	"desktop/internal/catalog/domain/derrors"
	"desktop/internal/catalog/domain/ports"
)

type UpdateCategoryUseCase struct {
	categoryRepository ports.CategoryRepository
}

func NewUpdateCategoryUseCase(categoryRepository ports.CategoryRepository) *UpdateCategoryUseCase {
	return &UpdateCategoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (uc *UpdateCategoryUseCase) Execute(
	ctx context.Context,
	organizationId string,
	cmd *commands.UpdateCategory,
) (*aggregates.Category, error) {
	existingCategory, err := uc.categoryRepository.FindOneById(ctx, organizationId, cmd.ID)

	if err != nil {
		return nil, err
	}

	if existingCategory == nil {
		return nil, derrors.ErrCategoryNotFound
	}

	if cmd.Name != "" && existingCategory.Name != cmd.Name {
		existingCategoryByName, err := uc.categoryRepository.FindOneByName(ctx, organizationId, cmd.Name)

		if err != nil {
			return nil, err
		}

		if existingCategoryByName != nil {
			return nil, derrors.ErrCategoryNameAlreadyExists
		}
	}

	if err := existingCategory.UpdateDetails(cmd.Name, cmd.Description); err != nil {
		return nil, err
	}

	if err := uc.categoryRepository.Update(ctx, organizationId, existingCategory); err != nil {
		return nil, err
	}

	return existingCategory, nil
}
