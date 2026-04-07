package usecases

import (
	"context"
	"desktop/internal/platform/catalog/domain/derrors"
	"desktop/internal/platform/catalog/domain/ports"
)

type DeleteProducts struct {
	productRepository ports.ProductRepository
}

func NewDeleteProducts(productRepository ports.ProductRepository) *DeleteProducts {
	return &DeleteProducts{
		productRepository: productRepository,
	}
}

func (uc *DeleteProducts) Execute(ctx context.Context, organizationId string, productId string) error {
	exisitingProduct, err := uc.productRepository.FindOneById(ctx, organizationId, productId)

	if err != nil {
		return err
	}

	if exisitingProduct == nil {
		return derrors.ErrProductNotFound
	}

	exisitingProduct.Delete()

	if err := uc.productRepository.Delete(ctx, exisitingProduct); err != nil {
		return err
	}

	return nil
}
