package responses

import (
	"desktop/internal/catalog/domain/aggregates"
	"desktop/internal/shared/db/platform"
)

type ProductResponse struct {
	ID             string  `json:"id"`
	Sku            string  `json:"sku"`
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	WholesalePrice int64   `json:"wholesalePrice"`
	SalePrice      int64   `json:"salePrice"`
	CostPrice      int64   `json:"costPrice"`
	Stock          int64   `json:"stock"`
	MinStock       int64   `json:"minStock"`
	OrganizationID string  `json:"organizationId"`
	CategoryID     *string `json:"categoryId"`
	SupplierID     *string `json:"supplierId"`
	CreatedAt      int64   `json:"createdAt"`
	UpdatedAt      int64   `json:"updatedAt"`
	DeletedAt      *int64  `json:"deletedAt"`
}

func ProductResponseFromDomain(product *aggregates.Product) *ProductResponse {
	return &ProductResponse{
		ID:             product.ID,
		Sku:            product.Sku.Value(),
		Name:           product.Name,
		Description:    product.Description,
		WholesalePrice: product.WholesalePrice.Amount(),
		SalePrice:      product.SalePrice.Amount(),
		CostPrice:      product.CostPrice.Amount(),
		Stock:          product.Stock,
		MinStock:       product.MinStock,
		OrganizationID: product.OrganizationID,
		CategoryID:     product.CategoryID,
		SupplierID:     product.SupplierID,
		CreatedAt:      platform.TimeToInt64(product.CreatedAt),
		UpdatedAt:      platform.TimeToInt64(product.UpdatedAt),
		DeletedAt:      platform.TimeToInt64Nullable(product.DeletedAt),
	}
}
