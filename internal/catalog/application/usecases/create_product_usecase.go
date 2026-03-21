package usecases

import (
	"context"
	"desktop/internal/catalog/application/commands"
	"desktop/internal/catalog/domain/aggregates"
	"desktop/internal/catalog/domain/derrors"
	"desktop/internal/catalog/domain/ports"
)

type CreateProductUseCase struct {
	productRepository ports.ProductRepository
}

func NewCreateProductUseCase(productRepository ports.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		productRepository: productRepository,
	}
}

func (u *CreateProductUseCase) Execute(
	ctx context.Context,
	organizationId string,
	command *commands.CreateProduct,
) (*aggregates.Product, error) {
	newProduct, errNewProduct := aggregates.NewProduct(
		command.Sku,
		command.Name,
		command.Description,
		command.WholesalePrice,
		command.SalePrice,
		command.CostPrice,
		command.Stock,
		command.MinStock,
		command.CategoryID,
		organizationId,
		command.SupplierID,
	)

	if errNewProduct != nil {
		return nil, errNewProduct
	}

	existisBySku, errExistisBySku := u.productRepository.FindOneBySku(ctx, organizationId, newProduct.Sku.Value())

	if errExistisBySku != nil {
		return nil, errExistisBySku
	}

	if existisBySku != nil {
		return nil, derrors.ErrProductSkuAlreadyExists
	}

	existisByName, errExistisByName := u.productRepository.FindOneByName(ctx, organizationId, newProduct.Name)

	if errExistisByName != nil {
		return nil, errExistisByName
	}

	if existisByName != nil {
		return nil, derrors.ErrProductNameAlreadyExists
	}

	if err := u.productRepository.Create(ctx, organizationId, newProduct); err != nil {
		return nil, err
	}

	return newProduct, nil
}
