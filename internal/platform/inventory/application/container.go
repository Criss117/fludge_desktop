package application

import (
	"desktop/internal/platform/inventory/application/usecases"
	"desktop/internal/platform/inventory/domain/ports"
)

type UseCasesContainer struct {
	CreateInventoryItem *usecases.CreateInventoryItem
	UpdateInventoryItem *usecases.UpdateInventoryItem
}

func NewUseCasesContainer(
	inventoryItemRepository ports.InventoryItemRepository,
) *UseCasesContainer {
	// InventoryItem - UseCases
	createInventoryItem := usecases.NewCreateInventoryItem(inventoryItemRepository)
	updateInventoryItem := usecases.NewUpdateInventoryItem(inventoryItemRepository, createInventoryItem)

	return &UseCasesContainer{
		CreateInventoryItem: createInventoryItem,
		UpdateInventoryItem: updateInventoryItem,
	}
}
