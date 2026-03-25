package mappers

import (
	"desktop/internal/platform/inventory/domain/aggregates"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

func MapInventoryItemFromDomain(inventoryItem *aggregates.InventoryItem) db.InventoryItem {
	return db.InventoryItem{
		ProductID:      inventoryItem.ProductID,
		OrganizationID: inventoryItem.OrganizationID,
		Stock:          inventoryItem.Stock,
		MinStock:       inventoryItem.MinStock,
		CreatedAt:      dbutils.TimeToInt64(inventoryItem.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(inventoryItem.UpdatedAt),
	}
}

func MapInventoryItemToDomain(inventoryItem db.InventoryItem) *aggregates.InventoryItem {
	return aggregates.ReconstituteInventoryItem(
		inventoryItem.ProductID,
		inventoryItem.OrganizationID,
		inventoryItem.Stock,
		inventoryItem.MinStock,
		dbutils.TimeFromInt64(inventoryItem.CreatedAt),
		dbutils.TimeFromInt64(inventoryItem.UpdatedAt),
	)
}
