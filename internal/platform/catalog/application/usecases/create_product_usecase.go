package usecases

import (
	"context"
	"desktop/internal/platform/catalog/application/commands"
	"desktop/internal/platform/catalog/domain/aggregates"
	"desktop/internal/platform/catalog/domain/derrors"
	"desktop/internal/platform/catalog/domain/ports"
	inventoryUsecases "desktop/internal/platform/inventory/application/usecases"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

type CreateProduct struct {
	txManager           dbutils.TxManager
	productRepository   ports.ProductRepository
	createInventoryItem inventoryUsecases.CreateInventoryItem
}

type CreateResponse struct {
	Product  *aggregates.Product
	Stock    int64
	MinStock int64
}

func NewCreateProduct(
	productRepository ports.ProductRepository,
	createInventoryItem inventoryUsecases.CreateInventoryItem,
) *CreateProduct {
	return &CreateProduct{
		productRepository:   productRepository,
		createInventoryItem: createInventoryItem,
	}
}

func (u *CreateProduct) Execute(
	ctx context.Context,
	organizationId string,
	cmd *commands.CreateProduct,
) (*CreateResponse, error) {
	exists, err := u.productRepository.ExistsBy(ctx, organizationId, ports.ExistsByParams{
		Name: cmd.Name,
		Sku:  cmd.Sku,
	})

	if err != nil {
		return nil, err
	}

	if exists.Name {
		return nil, derrors.ErrProductNameAlreadyExists
	}

	if exists.Sku {
		return nil, derrors.ErrProductSkuAlreadyExists
	}

	newProduct, err := aggregates.NewProduct(
		cmd.Sku,
		cmd.Name,
		cmd.Description,
		cmd.WholesalePrice,
		cmd.SalePrice,
		cmd.CostPrice,
		cmd.CategoryID,
		organizationId,
		cmd.SupplierID,
	)

	if err != nil {
		return nil, err
	}

	errTx := u.txManager.WithTx(ctx, func(q *db.Queries) error {
		if errDb := u.productRepository.Create(ctx, newProduct); errDb != nil {
			return errDb
		}

		if _, errDb := u.createInventoryItem.Execute(ctx, newProduct.ID, organizationId, cmd.Stock, cmd.MinStock); errDb != nil {
			return errDb
		}

		return nil
	})

	if errTx != nil {
		return nil, errTx
	}

	return &CreateResponse{
		Product:  newProduct,
		Stock:    cmd.Stock,
		MinStock: cmd.MinStock,
	}, nil
}
