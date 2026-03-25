package usecases

import (
	"context"
	"desktop/internal/platform/catalog/application/commands"
	"desktop/internal/platform/catalog/domain/ports"
)

type DeleteCategory struct {
	categoryrepository ports.CategoryRepository
}

func NewDeleteCategory(categoryrepository ports.CategoryRepository) *DeleteCategory {
	return &DeleteCategory{
		categoryrepository: categoryrepository,
	}
}

func (u *DeleteCategory) Execute(ctx context.Context, organizationId string, cmd *commands.DeleteManyCategories) error {
	if err := u.categoryrepository.DeleteMany(ctx, organizationId, cmd.IDs); err != nil {
		return err
	}

	return nil
}
