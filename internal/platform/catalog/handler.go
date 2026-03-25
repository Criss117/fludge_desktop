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
	app        *application.Container
	getCtx     GetCtxFunc
	getSession GetSessionFunc
}

func NewCatalogHandler(
	app *application.Container,
	getCtx GetCtxFunc,
	getSession GetSessionFunc,
) *CatalogHandler {
	return &CatalogHandler{
		getCtx:     getCtx,
		app:        app,
		getSession: getSession,
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

	categories, err := h.app.FindAllCategories.Execute(ctx, activeOrganizationId)

	if err != nil {
		return nil, err
	}

	categoryRes := make([]*responses.Category, len(categories))

	for i, category := range categories {
		categoryRes[i] = responses.CategoryResponseFromDomain(category)
	}

	return categoryRes, nil
}

func (h *CatalogHandler) DeleteCategory(cmd *commands.DeleteManyCategories) error {
	ctx := h.getCurrentContext()
	activeOrganizationId, err := h.getCurrentOrganizationId()

	if err != nil {
		return err
	}

	if err := h.app.DeleteCategory.Execute(ctx, activeOrganizationId, cmd); err != nil {
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

	updatedCategory, err := h.app.UpdateCategory.Execute(ctx, activeOrganizationId, cmd)

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

	createdCategory, err := h.app.CreateCategory.Execute(ctx, activeOrganizationId, cmd)

	if err != nil {
		return nil, err
	}

	return responses.CategoryResponseFromDomain(createdCategory), nil
}
