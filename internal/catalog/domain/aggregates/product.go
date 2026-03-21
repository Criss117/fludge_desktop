package aggregates

import (
	"desktop/internal/catalog/domain/derrors"
	"desktop/internal/catalog/domain/valueobjects"
	"desktop/internal/shared/lib"
	"time"
)

type Product struct {
	ID             string
	Sku            valueobjects.SKU
	Name           string
	Description    *string
	WholesalePrice valueobjects.Money
	SalePrice      valueobjects.Money
	CostPrice      valueobjects.Money
	Stock          int64
	MinStock       int64
	CategoryID     *string
	OrganizationID string
	SupplierID     *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

func NewProduct(
	sku, name string,
	description *string,
	wholesalePrice, salePrice, costPrice int64,
	stock, minStock int64,
	categoryID *string,
	organizationID string,
	supplierID *string,
) (*Product, error) {
	skuVO, err := valueobjects.NewSKU(sku)
	if err != nil {
		return nil, err
	}

	prices, err := valueobjects.NewPriceSet(costPrice, salePrice, wholesalePrice)
	if err != nil {
		return nil, err
	}

	if stock < 0 {
		return nil, derrors.ErrProductStockNegative
	}

	if minStock < 0 {
		minStock = 0
	}

	return &Product{
		ID:             lib.GenerateUUID(),
		Sku:            skuVO,
		Name:           name,
		Description:    description,
		WholesalePrice: prices.Wholesale,
		SalePrice:      prices.Sale,
		CostPrice:      prices.Cost,
		Stock:          stock,
		MinStock:       minStock,
		CategoryID:     categoryID,
		OrganizationID: organizationID,
		SupplierID:     supplierID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func ReconstituteProduct(
	id, sku, name string,
	description *string,
	wholesalePrice, salePrice, costPrice int64,
	stock, minStock int64,
	categoryID *string,
	organizationID string,
	supplierID *string,
	createdAt, updatedAt time.Time,
	deletedAt *time.Time,
) *Product {

	return &Product{
		ID:             id,
		Sku:            valueobjects.SKUFromStorage(sku),
		Name:           name,
		Description:    description,
		WholesalePrice: valueobjects.MoneyFromStorage(wholesalePrice),
		SalePrice:      valueobjects.MoneyFromStorage(salePrice),
		CostPrice:      valueobjects.MoneyFromStorage(costPrice),
		Stock:          stock,
		MinStock:       minStock,
		CategoryID:     categoryID,
		OrganizationID: organizationID,
		SupplierID:     supplierID,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		DeletedAt:      deletedAt,
	}
}

func (p *Product) UpdateDetails(name string, sku string, description *string) error {
	if name == "" {
		return derrors.ErrProductNameEmpty
	}

	skuVO, errSku := valueobjects.NewSKU(sku)
	if errSku != nil {
		return errSku
	}

	p.Name = name
	p.Sku = skuVO
	p.Description = description
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Product) UpdatePrices(costPrice, salePrice, wholesalePrice int64) error {
	prices, err := valueobjects.NewPriceSet(costPrice, salePrice, wholesalePrice)

	if err != nil {
		return err
	}

	p.CostPrice = prices.Cost
	p.SalePrice = prices.Sale
	p.WholesalePrice = prices.Wholesale
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Product) UpdateStock(stock, minStock int64) error {
	if minStock < 0 {
		minStock = 0
	}

	if stock < 0 {
		return derrors.ErrProductStockNegative
	}

	p.MinStock = minStock
	p.Stock = stock
	p.UpdatedAt = time.Now()

	return nil
}

func (p *Product) NeedsReorder() bool {
	// MinStock = 0 significa sin alerta de reorden
	if p.MinStock == 0 {
		return false
	}
	return p.Stock <= p.MinStock
}

func (p *Product) IsActive() bool             { return p.DeletedAt == nil }
func (p *Product) Equals(other *Product) bool { return p.ID == other.ID }

func (p *Product) Delete() {
	now := time.Now()
	p.DeletedAt = &now
	p.UpdatedAt = now
}
