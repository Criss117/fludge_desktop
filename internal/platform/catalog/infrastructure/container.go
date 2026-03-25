package infrastructure

import (
	"desktop/internal/platform/catalog/domain/ports"
	"desktop/internal/platform/catalog/infrastructure/repositories"
	"desktop/internal/shared/db"
)

type Container struct {
	CategoryRepository ports.CategoryRepository
	ProductRepository  ports.ProductRepository
}

func NewContainer(queries *db.Queries) *Container {
	categoryRepository := repositories.NewSqliteCategoryRepository(queries)
	productRepository := repositories.NewSqliteProductRepository(queries)

	return &Container{
		CategoryRepository: categoryRepository,
		ProductRepository:  productRepository,
	}
}
