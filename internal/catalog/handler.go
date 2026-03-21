package catalog

import (
	"context"
	"desktop/internal/catalog/application/commands"
	"desktop/internal/catalog/application/responses"
	"desktop/internal/catalog/application/usecases"
	"desktop/internal/catalog/domain/derrors"
	"desktop/internal/catalog/domain/ports"
	"desktop/internal/catalog/infrastructure/repositories"
	iamAggregates "desktop/internal/iam/domain/aggregates"
	"desktop/internal/shared/db"
)

type GetSessionFunc func() *iamAggregates.AppState

type CatalogHandler struct {
	ctx               context.Context
	productRepository ports.ProductRepository
	getSession        GetSessionFunc
}

func NewCatalogHandler(
	ctx context.Context,
	queries *db.Queries,
	getSession GetSessionFunc,

) *CatalogHandler {
	productRepository := repositories.NewSQLiteProductRepository(queries)

	return &CatalogHandler{
		ctx:               ctx,
		productRepository: productRepository,
		getSession:        getSession,
	}
}

func (h *CatalogHandler) activeOrganizationID() (string, error) {
	appState := h.getSession()
	if appState == nil {
		return "", derrors.ErrAppStateNotFound
	}

	if appState.ActiveOrganization == nil {
		return "", derrors.ErrNoActiveOrganization
	}

	return appState.ActiveOrganization.ID, nil
}

func (h *CatalogHandler) CreateProduct(createProductDto *commands.CreateProductCommand) (*responses.ProductResponse, error) {

	activeOrganizationId, err := h.activeOrganizationID()

	if err != nil {
		return nil, err
	}

	createProductUseCase := usecases.NewCreateProductUseCase(h.productRepository)

	newProduct, errCreating := createProductUseCase.Execute(h.ctx, activeOrganizationId, createProductDto)

	if errCreating != nil {
		return nil, errCreating
	}

	return responses.ProductResponseFromDomain(newProduct), nil
}

func (h *CatalogHandler) FindAllProducts() ([]*responses.ProductResponse, error) {
	activeOrganizationId, err := h.activeOrganizationID()

	if err != nil {
		return nil, err
	}

	findAllProductsUseCase := usecases.NewFindAllProductsUseCase(h.productRepository)

	products, err := findAllProductsUseCase.Execute(h.ctx, activeOrganizationId)

	if err != nil {
		return nil, err
	}

	responseProducts := make([]*responses.ProductResponse, len(products))

	for index, product := range products {
		responseProducts[index] = responses.ProductResponseFromDomain(product)
	}

	return responseProducts, nil
}
