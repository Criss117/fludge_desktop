package repositories

import (
	"context"
	"desktop/internal/platform/inventory/domain/aggregates"
	"desktop/internal/platform/inventory/infrastructure/mappers"
	"desktop/internal/shared/db"
	"desktop/internal/shared/db/dbutils"
)

type SqliteInventoryItemRepository struct {
	queries *db.Queries
}

func NewSqliteInventoryItemRepository(queries *db.Queries) *SqliteInventoryItemRepository {
	return &SqliteInventoryItemRepository{
		queries: queries,
	}
}

func (r *SqliteInventoryItemRepository) Create(ctx context.Context, inventoryItem *aggregates.InventoryItem) error {
	if errDb := r.queries.CreateInventoryItem(ctx, db.CreateInventoryItemParams{
		ProductID:      inventoryItem.ProductID,
		OrganizationID: inventoryItem.OrganizationID,
		Stock:          inventoryItem.Stock,
		MinStock:       inventoryItem.MinStock,
		CreatedAt:      dbutils.TimeToInt64(inventoryItem.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(inventoryItem.UpdatedAt),
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteInventoryItemRepository) Update(ctx context.Context, inventoryItem *aggregates.InventoryItem) error {
	if errDb := r.queries.UpdateInventoryItem(ctx, db.UpdateInventoryItemParams{
		ProductID:      inventoryItem.ProductID,
		OrganizationID: inventoryItem.OrganizationID,
		Stock:          inventoryItem.Stock,
		MinStock:       inventoryItem.MinStock,
		UpdatedAt:      dbutils.TimeToInt64(inventoryItem.UpdatedAt),
	}); errDb != nil {
		return errDb
	}

	return nil
}

func (r *SqliteInventoryItemRepository) FindOneByProductID(ctx context.Context, organizationID, productID string) (*aggregates.InventoryItem, error) {
	inventoryItem, err := r.queries.FindOneInventoryItem(ctx, db.FindOneInventoryItemParams{
		ProductID:      productID,
		OrganizationID: organizationID,
	})

	if err != nil {
		return nil, err
	}

	return mappers.MapInventoryItemToDomain(inventoryItem), nil
}
