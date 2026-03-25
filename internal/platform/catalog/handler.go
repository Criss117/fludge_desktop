package catalog

import (
	"context"
	"desktop/internal/appstate"
	"desktop/internal/platform/catalog/application"
	"desktop/internal/platform/catalog/application/commands"
	"desktop/internal/platform/catalog/application/responses"
)

type GetCtxFunc func() context.Context
type GetSessionFunc func() *appstate.SessionState

type CatalogHandler struct {
	usecases   *application.UseCasesContainer
	queries    *application.QueriesContainer
	getCtx     GetCtxFunc
	getSession GetSessionFunc
}

func NewCatalogHandler(
	usecases *application.UseCasesContainer,
	queries *application.QueriesContainer,
	getCtx GetCtxFunc,
	getSession GetSessionFunc,
) *CatalogHandler {
	return &CatalogHandler{
		getCtx:     getCtx,
		getSession: getSession,
		usecases:   usecases,
		queries:    queries,
	}
}

func (h *CatalogHandler) getCurrentContext() context.Context {
	ctx := context.Background()
	if h.getCtx != nil {
		ctx = h.getCtx()
	}
	return ctx
}

func (h *CatalogHandler) getCurrentOrganizationId() (string, error) {
	sessionState := h.getSession()

	if sessionState == nil {
		return "", nil
	}

	if sessionState.ActiveOrganization == nil {
		return "", nil
	}

	return sessionState.ActiveOrganization.ID, nil
}

func (h *CatalogHandler) FindAllCategories() ([]*responses.Category, error) {
	ctx := h.getCurrentContext()
	activeOrganizationId, err := h.getCurrentOrganizationId()

	if err != nil {
		return nil, err
	}

	categories, err := h.queries.FindAllCategories.Execute(ctx, activeOrganizationId)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (h *CatalogHandler) DeleteCategory(cmd *commands.DeleteManyCategories) error {
	ctx := h.getCurrentContext()
	activeOrganizationId, err := h.getCurrentOrganizationId()

	if err != nil {
		return err
	}

	if err := h.usecases.DeleteCategory.Execute(ctx, activeOrganizationId, cmd); err != nil {
		return err
	}

	return nil
}

func (h *CatalogHandler) UpdateCategory(cmd *commands.UpdateCategory) (*responses.Category, error) {
	ctx := h.getCurrentContext()
	activeOrganizationId, err := h.getCurrentOrganizationId()

	if err != nil {
		return nil, err
	}

	updatedCategory, err := h.usecases.UpdateCategory.Execute(ctx, activeOrganizationId, cmd)

	if err != nil {
		return nil, err
	}

	return responses.CategoryResponseFromDomain(updatedCategory), nil
}

func (h *CatalogHandler) CreateCategory(cmd *commands.CreateCategory) (*responses.Category, error) {
	ctx := h.getCurrentContext()
	activeOrganizationId, err := h.getCurrentOrganizationId()

	if err != nil {
		return nil, err
	}

	createdCategory, err := h.usecases.CreateCategory.Execute(ctx, activeOrganizationId, cmd)

	if err != nil {
		return nil, err
	}

	return responses.CategoryResponseFromDomain(createdCategory), nil
}

// Product -------------------------------------------------------------------------------------

func (h *CatalogHandler) FindAllProducts() ([]*responses.Product, error) {
	ctx := h.getCurrentContext()
	activeOrganizationId, err := h.getCurrentOrganizationId()

	if err != nil {
		return nil, err
	}

	products, err := h.queries.FindAllProducts.Execute(ctx, activeOrganizationId)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (h *CatalogHandler) CreateProduct(cmd *commands.CreateProduct) (*responses.Product, error) {
	ctx := h.getCurrentContext()
	activeOrganizationId, err := h.getCurrentOrganizationId()

	if err != nil {
		return nil, err
	}

	createdProduct, err := h.usecases.CreateProduct.Execute(ctx, activeOrganizationId, cmd)

	if err != nil {
		return nil, err
	}

	return responses.ProductFromDomain(createdProduct.Product, createdProduct.Stock, createdProduct.MinStock), nil
}

func (h *CatalogHandler) UpdateProduct(cmd *commands.UpdateProduct) (*responses.Product, error) {
	ctx := h.getCurrentContext()
	activeOrganizationId, err := h.getCurrentOrganizationId()

	if err != nil {
		return nil, err
	}

	product, err := h.usecases.UpdateProduct.Execute(ctx, activeOrganizationId, cmd)

	if err != nil {
		return nil, err
	}

	return responses.ProductFromDomain(product.Product, product.Stock, product.MinStock), nil
}
