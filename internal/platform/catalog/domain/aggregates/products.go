package aggregates

import (
	"desktop/internal/platform/catalog/domain/derrors"
	"desktop/internal/platform/catalog/domain/valueobjects"
	"desktop/internal/shared/events"
	"desktop/internal/shared/lib"
	"time"
)

type ProductCreated struct {
	events.BaseEvent
	Product *Product
}

type Product struct {
	ID             string
	Sku            valueobjects.SKU
	Name           string
	Description    *string
	WholesalePrice valueobjects.Money
	SalePrice      valueobjects.Money
	CostPrice      valueobjects.Money
	CategoryID     *string
	OrganizationID string
	SupplierID     *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	events         []events.DomainEvent
}

func NewProduct(
	sku, name string,
	description *string,
	wholesalePrice, salePrice, costPrice int64,
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

	return &Product{
		ID:             lib.GenerateUUID(),
		Sku:            skuVO,
		Name:           name,
		Description:    description,
		WholesalePrice: prices.Wholesale,
		SalePrice:      prices.Sale,
		CostPrice:      prices.Cost,
		CategoryID:     categoryID,
		OrganizationID: organizationID,
		SupplierID:     supplierID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func ReconstituteProduct(
	id, sku, name, organizationID string,
	description *string,
	wholesalePrice, salePrice, costPrice int64,
	categoryID, supplierID *string,
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

func (p *Product) IsActive() bool             { return p.DeletedAt == nil }
func (p *Product) Equals(other *Product) bool { return p.ID == other.ID }

func (p *Product) Delete() {
	now := time.Now()
	p.DeletedAt = &now
	p.UpdatedAt = now
}

func (p *Product) PullEvents() []events.DomainEvent {
	events := make([]events.DomainEvent, len(p.events))
	copy(events, p.events)
	p.events = nil
	return events
}

func (p *Product) record(event events.DomainEvent) {
	p.events = append(p.events, event)
}

func (p *Product) AfterCreate() {
	p.record(ProductCreated{
		BaseEvent: events.NewBaseEvent(events.ProductCreated, p.ID),
		Product:   p,
	})
}
