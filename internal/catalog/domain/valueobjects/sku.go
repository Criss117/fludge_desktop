// catalog/domain/valueobjects/sku.go
package valueobjects

import (
	"desktop/internal/catalog/domain/derrors"
	"strings"
)

type SKU struct {
	value string
}

func NewSKU(raw string) (SKU, error) {
	trimmed := strings.TrimSpace(raw)
	if len(trimmed) == 0 {
		return SKU{}, derrors.ErrSKUEmpty
	}
	return SKU{value: strings.ToUpper(trimmed)}, nil
}

func SKUFromStorage(raw string) SKU {
	return SKU{value: raw}
}

func (s SKU) Value() string         { return s.value }
func (s SKU) Equals(other SKU) bool { return s.value == other.value }
