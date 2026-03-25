package application

import (
	"desktop/internal/platform/catalog/application/usecases"
	"desktop/internal/platform/catalog/domain/ports"
	"desktop/internal/shared/db/dbutils"
)

type Container struct {
	FindAllCategories *usecases.FindAllCategories
	CreateCategory    *usecases.CreateCategory
	UpdateCategory    *usecases.UpdateCategory
	DeleteCategory    *usecases.DeleteCategory
}

func NewContainer(
	txManager dbutils.TxManager,
	categoryRepository ports.CategoryRepository,
	productRepository ports.ProductRepository,
) *Container {
	// Category - UseCases
	findAllCategories := usecases.NewFindAllCategories(categoryRepository)
	createCategory := usecases.NewCreateCategory(categoryRepository)
	updateCategory := usecases.NewUpdateCategory(categoryRepository)
	deleteCategory := usecases.NewDeleteCategory(categoryRepository)

	return &Container{
		FindAllCategories: findAllCategories,
		CreateCategory:    createCategory,
		UpdateCategory:    updateCategory,
		DeleteCategory:    deleteCategory,
	}
}
