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
	ctx                context.Context
	productRepository  ports.ProductRepository
	categoryRepository ports.CategoryRepository
	getSession         GetSessionFunc
}

func NewCatalogHandler(
	ctx context.Context,
	queries *db.Queries,
	getSession GetSessionFunc,
) *CatalogHandler {
	productRepository := repositories.NewSQLiteProductRepository(queries)
	categoryRepository := repositories.NewSQLiteCategoryRepository(queries)

	return &CatalogHandler{
		ctx:                ctx,
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
		getSession:         getSession,
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

func (h *CatalogHandler) UpdateProduct(
	updateProductDto *commands.UpdateProduct,
) (*responses.ProductResponse, error) {
	activeOrganizationId, err := h.activeOrganizationID()

	if err != nil {
		return nil, err
	}

	updateProductUseCase := usecases.NewUpdateProductUseCase(h.productRepository)

	newProduct, errUpdating := updateProductUseCase.Execute(h.ctx, activeOrganizationId, updateProductDto)

	if errUpdating != nil {
		return nil, errUpdating
	}

	return responses.ProductResponseFromDomain(newProduct), nil
}

func (h *CatalogHandler) CreateProduct(
	cmd *commands.CreateProduct,
) (*responses.ProductResponse, error) {
	activeOrganizationId, err := h.activeOrganizationID()

	if err != nil {
		return nil, err
	}

	createProductUseCase := usecases.NewCreateProductUseCase(h.productRepository)

	newProduct, errCreating := createProductUseCase.Execute(h.ctx, activeOrganizationId, cmd)

	if errCreating != nil {
		return nil, errCreating
	}

	return responses.ProductResponseFromDomain(newProduct), nil
}

// -----------------------------------------------------------------------------
// Categories
// -----------------------------------------------------------------------------

func (h *CatalogHandler) FindAllCategories() ([]*responses.CategoryResponse, error) {
	activeOrganizationId, err := h.activeOrganizationID()

	if err != nil {
		return nil, err
	}

	findAllCategoriesUseCase := usecases.NewFindAllCategoriesUseCase(h.categoryRepository)

	categories, err := findAllCategoriesUseCase.Execute(h.ctx, activeOrganizationId)

	if err != nil {
		return nil, err
	}

	responseCategories := make([]*responses.CategoryResponse, len(categories))

	for index, category := range categories {
		responseCategories[index] = responses.CategoryResponseFromDomain(category)
	}

	return responseCategories, nil
}

func (h *CatalogHandler) CreateCategory(
	cmd *commands.CreateCategory,
) (*responses.CategoryResponse, error) {
	activeOrganizationId, err := h.activeOrganizationID()

	if err != nil {
		return nil, err
	}

	createCategoryUseCase := usecases.NewCreateCategoryUseCase(h.categoryRepository)

	newCategory, errNewCategory := createCategoryUseCase.Execute(h.ctx, activeOrganizationId, cmd)

	if errNewCategory != nil {
		return nil, errNewCategory
	}

	return responses.CategoryResponseFromDomain(newCategory), nil
}

func (h *CatalogHandler) UpdateCategory(
	updateCategoryDto *commands.UpdateCategory,
) (*responses.CategoryResponse, error) {
	activeOrganizationId, err := h.activeOrganizationID()

	if err != nil {
		return nil, err
	}

	updateCategoryUseCase := usecases.NewUpdateCategoryUseCase(h.categoryRepository)

	newCategory, errUpdating := updateCategoryUseCase.Execute(h.ctx, activeOrganizationId, updateCategoryDto)

	if errUpdating != nil {
		return nil, errUpdating
	}

	return responses.CategoryResponseFromDomain(newCategory), nil
}

func (h *CatalogHandler) DeleteManyCategories(
	cmd *commands.DeleteManyCategories,
) error {
	activeOrganizationId, err := h.activeOrganizationID()

	if err != nil {
		return err
	}

	deleteManyCategoriesUseCase := usecases.NewDeleteCategoryUseCase(h.categoryRepository)

	return deleteManyCategoriesUseCase.Execute(h.ctx, activeOrganizationId, cmd)
}
