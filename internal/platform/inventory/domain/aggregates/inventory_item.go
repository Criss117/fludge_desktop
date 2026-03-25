package aggregates

import (
	"desktop/internal/platform/inventory/domain/derrors"
	"time"
)

type InventoryItem struct {
	ProductID      string
	OrganizationID string
	Stock          int64
	// MinStock is the minimum stock that can be stored in the inventory,
	// it can be negative to allow stocks to be less than min stock
	MinStock  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewInventoryItem(
	productID string,
	organizationID string,
	stock, minStock int64,
) (*InventoryItem, error) {
	if minStock >= 0 && stock < minStock {
		return nil, derrors.ErrStockGreaterThanMin
	}

	now := time.Now()

	return &InventoryItem{
		ProductID:      productID,
		OrganizationID: organizationID,
		Stock:          stock,
		MinStock:       minStock,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func ReconstituteInventoryItem(
	productID, organizationID string,
	stock, minStock int64,
	createdAt, updatedAt time.Time,
) *InventoryItem {
	return &InventoryItem{
		ProductID:      productID,
		OrganizationID: organizationID,
		Stock:          stock,
		MinStock:       minStock,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}

func (i *InventoryItem) AllowNegativeStock() bool { return i.MinStock < 0 }

func (i *InventoryItem) Update(stock, minStock int64) error {
	if !i.AllowNegativeStock() {
		if stock < minStock {
			return derrors.ErrStockGreaterThanMin
		}
	}

	if minStock < -1 {
		return derrors.ErrInvalidMinStock
	}

	i.Stock = stock
	i.MinStock = minStock
	i.UpdatedAt = time.Now()
	return nil
}

func (i *InventoryItem) IncreaseStock(amount int64) error {
	if amount <= 0 {
		return derrors.ErrInvalidQuantity
	}

	i.Stock += amount
	i.UpdatedAt = time.Now()

	return nil
}

func (i *InventoryItem) DecreaseStock(amount int64) error {
	if amount <= 0 {
		return derrors.ErrInvalidQuantity
	}

	if !i.AllowNegativeStock() {
		if i.Stock-amount < i.MinStock {
			return derrors.ErrStockGreaterThanMin
		}
	}

	i.Stock -= amount
	i.UpdatedAt = time.Now()

	return nil
}

func (i *InventoryItem) IsStockSufficient(amount int64) bool {
	if amount <= 0 {
		return false
	}

	if i.AllowNegativeStock() {
		return true
	}

	return i.Stock-amount >= i.MinStock
}

func (i *InventoryItem) UpdateStock(stock int64) error {
	if !i.AllowNegativeStock() {
		if stock < i.MinStock {
			return derrors.ErrStockGreaterThanMin
		}
	}

	i.Stock = stock
	i.UpdatedAt = time.Now()
	return nil
}
