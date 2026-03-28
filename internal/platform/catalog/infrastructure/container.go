package infrastructure

import (
	"desktop/internal/platform/catalog/application"
	"desktop/internal/platform/catalog/domain/ports"
	"desktop/internal/platform/catalog/infrastructure/handlers"
	"desktop/internal/platform/catalog/infrastructure/repositories"
	"desktop/internal/shared/db"
	"desktop/internal/shared/infrastructure"
)

type CatalogHandlerContainer struct {
	ProductHandler  handlers.CatalogProductHandler
	CategoryHandler handlers.CatalogCategoryHandler
}

type CatalogRepositoryContainer struct {
	ProductRepository  ports.ProductRepository
	CategoryRepository ports.CategoryRepository
}

func NewRepositoryContainer(querires *db.Queries) *CatalogRepositoryContainer {
	productRepository := repositories.NewSqliteProductRepository(querires)
	categoryRepository := repositories.NewSqliteCategoryRepository(querires)

	return &CatalogRepositoryContainer{
		ProductRepository:  productRepository,
		CategoryRepository: categoryRepository,
	}
}

func NewHandlerContainer(
	usecases *application.UseCasesContainer,
	queries *application.QueriesContainer,
	getCtx infrastructure.GetCtxFunc,
	getSession infrastructure.GetSessionFunc,
) CatalogHandlerContainer {
	productHandler := handlers.NewCatalogProductHandler(
		usecases,
		queries,
		getCtx,
		getSession,
	)

	categoryHandler := handlers.NewCatalogCategoryHandler(
		usecases,
		queries,
		getCtx,
		getSession,
	)

	return CatalogHandlerContainer{
		ProductHandler:  productHandler,
		CategoryHandler: categoryHandler,
	}
}
