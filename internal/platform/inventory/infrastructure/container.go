package infrastructure

import (
	"desktop/internal/platform/inventory/domain/ports"
	"desktop/internal/platform/inventory/infrastructure/repositories"
	"desktop/internal/shared/db"
)

type RepositoriesContainer struct {
	InventoryItemRepository ports.InventoryItemRepository
}

func NewRepositoriesContainer(queries *db.Queries) *RepositoriesContainer {
	// InventoryItem - Repositories
	inventoryItemRepository := repositories.NewSqliteInventoryItemRepository(queries)

	return &RepositoriesContainer{
		InventoryItemRepository: inventoryItemRepository,
	}
}
