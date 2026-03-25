package ports

import (
	"context"
	"desktop/internal/platform/inventory/domain/aggregates"
)

type InventoryItemRepository interface {
	Create(ctx context.Context, inventoryItem *aggregates.InventoryItem) error
	Update(ctx context.Context, inventoryItem *aggregates.InventoryItem) error

	FindOneByProductID(ctx context.Context, productID, organizationID string) (*aggregates.InventoryItem, error)
}
