package catalog

import (
	"context"
	"desktop/internal/catalog/application/commands"
	"desktop/internal/catalog/application/responses"
	"desktop/internal/catalog/application/usecases"
	"desktop/internal/catalog/domain/ports"
	"desktop/internal/catalog/infrastructure/repositories"
	"desktop/internal/shared/db"
)

type CatalogHandler struct {
	ctx               context.Context
	productRepository ports.ProductRepository
}

func NewCatalogHandler(
	ctx context.Context,
	queries *db.Queries,
) *CatalogHandler {

	productRepository := repositories.NewSQLiteProductRepository(queries)

	return &CatalogHandler{
		ctx:               ctx,
		productRepository: productRepository,
	}
}

func (h *CatalogHandler) CreateProduct(createProductDto *commands.CreateProductCommand) (*responses.ProductResponse, error) {
	createProductUseCase := usecases.NewCreateProductUseCase(h.productRepository)

	newProduct, err := createProductUseCase.Execute(h.ctx, createProductDto)

	if err != nil {
		return nil, err
	}

	return responses.ProductResponseFromDomain(newProduct), nil
}
