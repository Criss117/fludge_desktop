package usecases

import (
	"context"
	"desktop/internal/platform/inventory/domain/aggregates"
	"desktop/internal/platform/inventory/domain/ports"
)

type CreateInventoryItem struct {
	inventoryItemRepository ports.InventoryItemRepository
}

func NewCreateInventoryItem(inventoryItemRepository ports.InventoryItemRepository) *CreateInventoryItem {
	return &CreateInventoryItem{
		inventoryItemRepository: inventoryItemRepository,
	}
}

func (u *CreateInventoryItem) Execute(
	ctx context.Context,
	productID, organizationID string,
	stock, minStock int64,
) error {
	newInvItem, err := aggregates.NewInventoryItem(
		productID,
		organizationID,
		stock,
		minStock,
	)

	if err != nil {
		return err
	}

	return u.inventoryItemRepository.Create(ctx, newInvItem)
}
