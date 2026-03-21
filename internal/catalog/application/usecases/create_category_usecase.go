package usecases

import (
	"context"
	"desktop/internal/catalog/application/commands"
	"desktop/internal/catalog/domain/aggregates"
	"desktop/internal/catalog/domain/derrors"
	"desktop/internal/catalog/domain/ports"
)

type CreateCategoryUseCase struct {
	categoryRepository ports.CategoryRepository
}

func NewCreateCategoryUseCase(categoryRepository ports.CategoryRepository) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (u *CreateCategoryUseCase) Execute(
	ctx context.Context,
	organizationId string,
	command *commands.CreateCategoryCommand,
) (*aggregates.Category, error) {
	newCategory, errNewCategory := aggregates.NewCategory(
		command.Name,
		command.Description,
		organizationId,
	)

	if errNewCategory != nil {
		return nil, errNewCategory
	}

	existingCategoryByName, errExistingCategoryByName := u.categoryRepository.FindOneByName(
		ctx,
		organizationId,
		newCategory.Name,
	)

	if errExistingCategoryByName != nil {
		return nil, errExistingCategoryByName
	}

	if existingCategoryByName != nil {
		return nil, derrors.ErrCategoryNameAlreadyExists
	}

	if err := u.categoryRepository.Create(ctx, organizationId, newCategory); err != nil {
		return nil, err
	}

	return newCategory, nil
}
