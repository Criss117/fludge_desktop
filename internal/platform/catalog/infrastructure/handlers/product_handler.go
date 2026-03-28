package handlers

import (
	"desktop/internal/platform/catalog/application"
	"desktop/internal/platform/catalog/application/commands"
	"desktop/internal/platform/catalog/application/responses"
	"desktop/internal/shared/infrastructure"
)

type CatalogProductHandler struct {
	baseHandler infrastructure.BaseHandler
	usecases    *application.UseCasesContainer
	queries     *application.QueriesContainer
}

func NewCatalogProductHandler(
	usecases *application.UseCasesContainer,
	queries *application.QueriesContainer,
	getCtx infrastructure.GetCtxFunc,
	getSession infrastructure.GetSessionFunc,
) CatalogProductHandler {
	return CatalogProductHandler{
		usecases:    usecases,
		queries:     queries,
		baseHandler: infrastructure.NewBaseHandler(getCtx, getSession),
	}
}

func (h *CatalogProductHandler) FindAllProducts() ([]*responses.Product, error) {
	ctx := h.baseHandler.Context()
	activeOrganization, err := h.baseHandler.CurrentOrganization()

	if err != nil {
		return nil, err
	}

	products, err := h.queries.FindAllProducts.Execute(ctx, activeOrganization.ID)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (h *CatalogProductHandler) CreateProduct(cmd *commands.CreateProduct) (*responses.Product, error) {
	ctx := h.baseHandler.Context()
	activeOrganization, err := h.baseHandler.CurrentOrganization()

	if err != nil {
		return nil, err
	}

	createdProduct, err := h.usecases.CreateProduct.Execute(ctx, activeOrganization.ID, cmd)

	if err != nil {
		return nil, err
	}

	return responses.ProductFromDomain(createdProduct.Product, createdProduct.Stock, createdProduct.MinStock), nil
}

func (h *CatalogProductHandler) UpdateProduct(cmd *commands.UpdateProduct) (*responses.Product, error) {
	ctx := h.baseHandler.Context()
	activeOrganization, err := h.baseHandler.CurrentOrganization()

	if err != nil {
		return nil, err
	}

	product, err := h.usecases.UpdateProduct.Execute(ctx, activeOrganization.ID, cmd)

	if err != nil {
		return nil, err
	}

	return responses.ProductFromDomain(product.Product, product.Stock, product.MinStock), nil
}
