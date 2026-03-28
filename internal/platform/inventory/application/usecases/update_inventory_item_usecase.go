package usecases

import (
	"context"
	"desktop/internal/platform/inventory/domain/aggregates"
	"desktop/internal/platform/inventory/domain/ports"
)

type UpdateInventoryItem struct {
	inventoryItemRepository ports.InventoryItemRepository
	createInventoryItem     *CreateInventoryItem
}

func NewUpdateInventoryItem(
	inventoryItemRepository ports.InventoryItemRepository,
	createInventoryItem *CreateInventoryItem,
) *UpdateInventoryItem {
	return &UpdateInventoryItem{
		inventoryItemRepository: inventoryItemRepository,
		createInventoryItem:     createInventoryItem,
	}
}

func (u *UpdateInventoryItem) Execute(
	ctx context.Context,
	organizationId string,
	productId string,
	stock int64,
	minStock int64,
) (*aggregates.InventoryItem, error) {
	inventoryItem, err := u.inventoryItemRepository.FindOneByProductID(ctx, organizationId, productId)

	if err != nil {
		return nil, err
	}

	if inventoryItem == nil {

		newItem, err := u.createInventoryItem.Execute(ctx, productId, organizationId, stock, minStock)

		if err != nil {
			return nil, err
		}

		return newItem, nil
	}

	errUpdating := inventoryItem.Update(stock, minStock)

	if errUpdating != nil {
		return nil, errUpdating
	}

	if err := u.inventoryItemRepository.Update(ctx, inventoryItem); err != nil {
		return nil, err
	}

	return inventoryItem, nil
}
