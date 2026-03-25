package application

import (
	"desktop/internal/platform/inventory/application/usecases"
	"desktop/internal/platform/inventory/domain/ports"
)

type UseCasesContainer struct {
	CreateInventoryItem *usecases.CreateInventoryItem
}

func NewUseCasesContainer(
	inventoryItemRepository ports.InventoryItemRepository,
) *UseCasesContainer {
	// InventoryItem - UseCases
	createInventoryItem := usecases.NewCreateInventoryItem(inventoryItemRepository)

	return &UseCasesContainer{
		CreateInventoryItem: createInventoryItem,
	}
}
