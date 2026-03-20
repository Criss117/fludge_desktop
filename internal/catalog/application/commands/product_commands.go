package commands

type CreateProductCommand struct {
	Name           string  `json:"name"`
	Sku            string  `json:"sku"`
	Description    *string `json:"description"`
	WholesalePrice int64   `json:"wholesalePrice"`
	SalePrice      int64   `json:"salePrice"`
	CostPrice      int64   `json:"costPrice"`
	Stock          int64   `json:"stock"`
	MinStock       int64   `json:"minStock"`
	OrganizationID string  `json:"organizationId"`
	CategoryID     *string `json:"categoryId"`
	SupplierID     *string `json:"supplierId"`
}
