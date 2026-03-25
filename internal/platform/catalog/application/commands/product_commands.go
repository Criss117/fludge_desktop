package commands

type CreateProduct struct {
	Name           string  `json:"name"`
	Sku            string  `json:"sku"`
	Description    *string `json:"description"`
	WholesalePrice int64   `json:"wholesalePrice"`
	SalePrice      int64   `json:"salePrice"`
	CostPrice      int64   `json:"costPrice"`
	Stock          int64   `json:"stock"`
	MinStock       int64   `json:"minStock"`
	CategoryID     *string `json:"categoryId"`
	SupplierID     *string `json:"supplierId"`
}

type UpdateProduct struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Sku            string  `json:"sku"`
	Description    *string `json:"description"`
	WholesalePrice int64   `json:"wholesalePrice"`
	SalePrice      int64   `json:"salePrice"`
	CostPrice      int64   `json:"costPrice"`
	Stock          int64   `json:"stock"`
	MinStock       int64   `json:"minStock"`
	CategoryID     *string `json:"categoryId"`
	SupplierID     *string `json:"supplierId"`
}
