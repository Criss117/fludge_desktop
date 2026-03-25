package application

import (
	"desktop/internal/platform/catalog/application/queries"
	"desktop/internal/platform/catalog/application/usecases"
	inventoryUsecases "desktop/internal/platform/inventory/application/usecases"

	"desktop/internal/platform/catalog/domain/ports"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

type UseCasesContainer struct {
	CreateCategory *usecases.CreateCategory
	UpdateCategory *usecases.UpdateCategory
	DeleteCategory *usecases.DeleteCategory
	CreateProduct  *usecases.CreateProduct
}

type QueriesContainer struct {
	FindAllCategories *queries.FindAllCategories
}

func NewUseCasesContainer(
	txManager dbutils.TxManager,
	categoryRepository ports.CategoryRepository,
	productRepository ports.ProductRepository,
	createInventoryItem inventoryUsecases.CreateInventoryItem,
) *UseCasesContainer {
	// Category - UseCases
	createCategory := usecases.NewCreateCategory(categoryRepository)
	updateCategory := usecases.NewUpdateCategory(categoryRepository)
	deleteCategory := usecases.NewDeleteCategory(categoryRepository)

	// Product - UseCases
	createProduct := usecases.NewCreateProduct(productRepository, createInventoryItem)

	return &UseCasesContainer{
		CreateCategory: createCategory,
		UpdateCategory: updateCategory,
		DeleteCategory: deleteCategory,
		CreateProduct:  createProduct,
	}
}

func NewQueriesContainer(dbQueries *db.Queries) *QueriesContainer {
	// Category - Queries
	findAllCategories := queries.NewFindAllCategories(dbQueries)

	return &QueriesContainer{
		FindAllCategories: findAllCategories,
	}
}
