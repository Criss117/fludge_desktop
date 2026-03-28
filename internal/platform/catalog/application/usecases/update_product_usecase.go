package usecases

import (
	"context"
	"desktop/internal/platform/catalog/application/commands"
	"desktop/internal/platform/catalog/domain/aggregates"
	"desktop/internal/platform/catalog/domain/derrors"
	"desktop/internal/platform/catalog/domain/ports"
	inventoryUsecases "desktop/internal/platform/inventory/application/usecases"
	"desktop/internal/shared/db/dbutils"
	"log"
)

type UpdateProduct struct {
	txManager           dbutils.TxManager
	productRepository   ports.ProductRepository
	updateInventoryItem inventoryUsecases.UpdateInventoryItem
}

type UpdateResponse struct {
	Product  *aggregates.Product
	Stock    int64
	MinStock int64
}

func NewUpdateProduct(productRepository ports.ProductRepository, updateInventoryItem inventoryUsecases.UpdateInventoryItem) *UpdateProduct {
	return &UpdateProduct{
		productRepository:   productRepository,
		updateInventoryItem: updateInventoryItem,
	}
}

func (u *UpdateProduct) Execute(
	ctx context.Context,
	organizationId string,
	cmd *commands.UpdateProduct,
) (*UpdateResponse, error) {
	existingProduct, err := u.productRepository.FindOneById(ctx, organizationId, cmd.ID)

	if err != nil {
		return nil, err
	}

	if existingProduct == nil {
		return nil, derrors.ErrProductNotFound
	}

	if existingProduct.Name != cmd.Name || existingProduct.Sku.Value() != cmd.Sku {
		exists, err := u.productRepository.ExistsBy(ctx, organizationId, ports.ExistsByParams{
			Name: cmd.Name,
			Sku:  cmd.Sku,
		})

		if err != nil {
			return nil, err
		}

		if existingProduct.Name != cmd.Name && exists.Name {
			return nil, derrors.ErrProductNameAlreadyExists
		}

		if existingProduct.Sku.Value() != cmd.Sku && exists.Sku {
			return nil, derrors.ErrProductSkuAlreadyExists
		}
	}

	existingProduct.UpdateDetails(
		cmd.Name,
		cmd.Sku,
		cmd.Description,
	)

	existingProduct.UpdatePrices(
		cmd.CostPrice,
		cmd.SalePrice,
		cmd.WholesalePrice,
	)

	if errDb := u.productRepository.Update(ctx, existingProduct); errDb != nil {
		return nil, errDb
	}

	log.Println("Despues de actualizar product")

	if _, errDb := u.updateInventoryItem.Execute(ctx, existingProduct.ID, organizationId, cmd.Stock, cmd.MinStock); errDb != nil {
		return nil, errDb
	}

	log.Println("Despues de actualizar item")
	// errTx := u.txManager.WithTx(ctx, func(q *db.Queries) error {

	// 	return nil
	// })

	// if errTx != nil {
	// 	return nil, errTx
	// }

	return &UpdateResponse{
		Product:  existingProduct,
		Stock:    cmd.Stock,
		MinStock: cmd.MinStock,
	}, nil
}
