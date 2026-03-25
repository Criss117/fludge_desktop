package derrors

import "errors"

var (
	ErrStockNegative       = errors.New("El stock no puede ser negativo")
	ErrMinStockNegative    = errors.New("El stock mínimo no puede ser negativo")
	ErrStockGreaterThanMin = errors.New("El stock no puede ser mayor que el stock mínimo")
	ErrInvalidMinStock     = errors.New("El stock mínimo no puede ser menor a -1")
	ErrInvalidQuantity     = errors.New("La cantidad no puede ser menor a cero")
)
