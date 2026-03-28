package usecases

import (
	"context"
	"desktop/internal/platform/inventory/domain/aggregates"
	"desktop/internal/platform/inventory/domain/ports"
	"log"
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

	log.Println("Execute 1")

	if inventoryItem == nil {
		log.Println("Execute 1-1")

		newItem, err := u.createInventoryItem.Execute(ctx, productId, organizationId, stock, minStock)

		log.Println("Execute 1-2")

		if err != nil {
			return nil, err
		}

		log.Println("Execute 1-3")

		return newItem, nil
	}

	log.Println("Execute 2")

	errUpdating := inventoryItem.Update(stock, minStock)

	if errUpdating != nil {
		return nil, errUpdating
	}

	if err := u.inventoryItemRepository.Update(ctx, inventoryItem); err != nil {
		return nil, err
	}

	return inventoryItem, nil
}
