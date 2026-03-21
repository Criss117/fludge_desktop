package usecases

import (
	"context"
	"desktop/internal/catalog/domain/aggregates"
	"desktop/internal/catalog/domain/ports"
)

type FindAllProductsUseCase struct {
	productRepository ports.ProductRepository
}

func NewFindAllProductsUseCase(productRepository ports.ProductRepository) *FindAllProductsUseCase {
	return &FindAllProductsUseCase{
		productRepository: productRepository,
	}
}

func (u *FindAllProductsUseCase) Execute(ctx context.Context, organizationId string) ([]*aggregates.Product, error) {
	return u.productRepository.FindAllProducts(ctx, organizationId)
}
