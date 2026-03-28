package handlers

import (
	"desktop/internal/platform/catalog/application"
	"desktop/internal/platform/catalog/application/commands"
	"desktop/internal/platform/catalog/application/responses"
	"desktop/internal/shared/infrastructure"
)

type CatalogCategoryHandler struct {
	baseHandler infrastructure.BaseHandler
	usecases    *application.UseCasesContainer
	queries     *application.QueriesContainer
}

func NewCatalogCategoryHandler(
	usecases *application.UseCasesContainer,
	queries *application.QueriesContainer,
	getCtx infrastructure.GetCtxFunc,
	getSession infrastructure.GetSessionFunc,
) CatalogCategoryHandler {
	return CatalogCategoryHandler{
		usecases:    usecases,
		queries:     queries,
		baseHandler: infrastructure.NewBaseHandler(getCtx, getSession),
	}
}

func (h *CatalogCategoryHandler) FindAllCategories() ([]*responses.Category, error) {
	ctx := h.baseHandler.Context()
	activeOrganization, err := h.baseHandler.CurrentOrganization()

	if err != nil {
		return nil, err
	}

	categories, err := h.queries.FindAllCategories.Execute(ctx, activeOrganization.ID)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (h *CatalogCategoryHandler) DeleteManyCategories(cmd *commands.DeleteManyCategories) error {
	ctx := h.baseHandler.Context()
	activeOrganization, err := h.baseHandler.CurrentOrganization()

	if err != nil {
		return err
	}

	if err := h.usecases.DeleteCategory.Execute(ctx, activeOrganization.ID, cmd); err != nil {
		return err
	}

	return nil
}

func (h *CatalogCategoryHandler) UpdateCategory(cmd *commands.UpdateCategory) (*responses.Category, error) {
	ctx := h.baseHandler.Context()
	activeOrganization, err := h.baseHandler.CurrentOrganization()

	if err != nil {
		return nil, err
	}

	updatedCategory, err := h.usecases.UpdateCategory.Execute(ctx, activeOrganization.ID, cmd)

	if err != nil {
		return nil, err
	}

	return responses.CategoryResponseFromDomain(updatedCategory), nil
}

func (h *CatalogCategoryHandler) CreateCategory(cmd *commands.CreateCategory) (*responses.Category, error) {
	ctx := h.baseHandler.Context()
	activeOrganization, err := h.baseHandler.CurrentOrganization()

	if err != nil {
		return nil, err
	}

	createdCategory, err := h.usecases.CreateCategory.Execute(ctx, activeOrganization.ID, cmd)

	if err != nil {
		return nil, err
	}

	return responses.CategoryResponseFromDomain(createdCategory), nil
}
