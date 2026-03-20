package derrors

import "errors"

var (
	ErrSalePriceBelowCost       = errors.New("El precio de venta debe ser mayor o igual que el precio de costo")
	ErrMoneyNegative            = errors.New("El monto no puede ser negativo")
	ErrWholesalePriceAboveSale  = errors.New("El precio de venta debe ser menor o igual que el precio de venta")
	ErrSKUEmpty                 = errors.New("El SKU no puede estar vacío")
	ErrProductStockNegative     = errors.New("El stock no puede ser negativo")
	ErrProductMinStockNegative  = errors.New("El stock mínimo no puede ser negativo")
	ErrProductSkuAlreadyExists  = errors.New("El SKU ya esta en uso")
	ErrProductNameAlreadyExists = errors.New("El nombre ya esta en uso")
	ErrProductNotFound          = errors.New("El producto no existe")
)
