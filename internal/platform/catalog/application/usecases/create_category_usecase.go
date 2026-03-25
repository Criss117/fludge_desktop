package usecases

import (
	"context"
	"desktop/internal/platform/catalog/application/commands"
	"desktop/internal/platform/catalog/domain/aggregates"
	"desktop/internal/platform/catalog/domain/derrors"
	"desktop/internal/platform/catalog/domain/ports"
)

type CreateCategory struct {
	categoryrepository ports.CategoryRepository
}

func NewCreateCategory(categoryrepository ports.CategoryRepository) *CreateCategory {
	return &CreateCategory{
		categoryrepository: categoryrepository,
	}
}

func (u *CreateCategory) Execute(ctx context.Context, organizationId string, cmd *commands.CreateCategory) (*aggregates.Category, error) {
	exisitngCategory, errRepo := u.categoryrepository.FindOneByName(ctx, organizationId, cmd.Name)

	if errRepo != nil {
		return nil, errRepo
	}

	if exisitngCategory != nil {
		return nil, derrors.ErrCategoryNameAlreadyExists
	}

	category, errAgg := aggregates.NewCategory(
		cmd.Name,
		cmd.Description,
		organizationId,
	)

	if errAgg != nil {
		return nil, errAgg
	}

	if err := u.categoryrepository.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}
