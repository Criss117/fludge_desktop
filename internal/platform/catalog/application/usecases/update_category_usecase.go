package usecases

import (
	"context"
	"desktop/internal/platform/catalog/application/commands"
	"desktop/internal/platform/catalog/domain/aggregates"
	"desktop/internal/platform/catalog/domain/derrors"
	"desktop/internal/platform/catalog/domain/ports"
)

type UpdateCategory struct {
	categoryrepository ports.CategoryRepository
}

func NewUpdateCategory(categoryrepository ports.CategoryRepository) *UpdateCategory {
	return &UpdateCategory{
		categoryrepository: categoryrepository,
	}
}

func (u *UpdateCategory) Execute(ctx context.Context, organizationId string, cmd *commands.UpdateCategory) (*aggregates.Category, error) {
	exisitngCategory, errRepo := u.categoryrepository.FindOneById(ctx, organizationId, cmd.ID)

	if errRepo != nil {
		return nil, errRepo
	}

	if exisitngCategory == nil {
		return nil, derrors.ErrCategoryNotFound
	}

	category, errAgg := aggregates.NewCategory(
		exisitngCategory.Name,
		exisitngCategory.Description,
		organizationId,
	)

	if errAgg != nil {
		return nil, errAgg
	}

	if err := u.categoryrepository.Update(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}
