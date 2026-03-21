package usecases

import (
	"context"
	"desktop/internal/catalog/application/commands"
	"desktop/internal/catalog/domain/aggregates"
	"desktop/internal/catalog/domain/derrors"
	"desktop/internal/catalog/domain/ports"
)

type UpdateProductUseCase struct {
	productRepository ports.ProductRepository
}

func NewUpdateProductUseCase(productRepository ports.ProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productRepository: productRepository,
	}
}

func (u *UpdateProductUseCase) Execute(
	ctx context.Context,
	organizationId string,
	command *commands.UpdateProduct,
) (*aggregates.Product, error) {
	existingProduct, err := u.productRepository.FindOneById(ctx, organizationId, command.ID)

	if err != nil {
		return nil, err
	}

	if existingProduct == nil {
		return nil, derrors.ErrProductNotFound
	}

	if command.Sku != "" && existingProduct.Sku.Value() != command.Sku {
		existingProductBySku, err := u.productRepository.FindOneBySku(ctx, organizationId, command.Sku)

		if err != nil {
			return nil, err
		}

		if existingProductBySku != nil {
			return nil, derrors.ErrProductSkuAlreadyExists
		}
	}

	if command.Name != "" && existingProduct.Name != command.Name {
		existingProductByName, err := u.productRepository.FindOneByName(ctx, organizationId, command.Name)

		if err != nil {
			return nil, err
		}

		if existingProductByName != nil {
			return nil, derrors.ErrProductNameAlreadyExists
		}
	}

	if err := existingProduct.UpdateDetails(
		command.Name,
		command.Sku,
		command.Description,
	); err != nil {
		return nil, err
	}

	if err := existingProduct.UpdatePrices(
		command.CostPrice,
		command.SalePrice,
		command.WholesalePrice,
	); err != nil {
		return nil, err
	}

	if err := existingProduct.UpdateStock(command.Stock, command.MinStock); err != nil {
		return nil, err
	}

	if err := u.productRepository.Update(ctx, organizationId, existingProduct); err != nil {
		return nil, err
	}

	return existingProduct, nil
}
