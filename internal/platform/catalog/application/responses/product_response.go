package responses

import (
	"desktop/internal/platform/catalog/domain/aggregates"
	"desktop/internal/shared/db/dbutils"
)

type Product struct {
	ID             string  `json:"id"`
	Sku            string  `json:"sku"`
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	WholesalePrice int64   `json:"wholesalePrice"`
	SalePrice      int64   `json:"salePrice"`
	CostPrice      int64   `json:"costPrice"`
	CategoryID     *string `json:"categoryId"`
	OrganizationID string  `json:"organizationId"`
	SupplierID     *string `json:"supplierId"`
	CreatedAt      int64   `json:"createdAt"`
	UpdatedAt      int64   `json:"updatedAt"`
	DeletedAt      *int64  `json:"deletedAt"`
	Stock          int64   `json:"stock"`
	MinStock       int64   `json:"minStock"`
}

func ProductFromDomain(product *aggregates.Product, stock int64, minStock int64) *Product {
	if product == nil {
		return nil
	}

	return &Product{
		ID:             product.ID,
		Sku:            product.Sku.Value(),
		Name:           product.Name,
		Description:    product.Description,
		WholesalePrice: product.WholesalePrice.Amount(),
		SalePrice:      product.SalePrice.Amount(),
		CostPrice:      product.CostPrice.Amount(),
		CategoryID:     product.CategoryID,
		OrganizationID: product.OrganizationID,
		SupplierID:     product.SupplierID,
		CreatedAt:      dbutils.TimeToInt64(product.CreatedAt),
		UpdatedAt:      dbutils.TimeToInt64(product.UpdatedAt),
		DeletedAt:      dbutils.TimeToInt64Nullable(product.DeletedAt),
	}
}
