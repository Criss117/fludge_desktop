// catalog/domain/valueobjects/money.go
package valueobjects

import "desktop/internal/catalog/domain/derrors"

// Money representa precios en centavos — evita problemas de punto flotante
type Money struct {
	amount int64
}

func NewMoney(amount int64) (Money, error) {
	if amount < 0 {
		return Money{}, derrors.ErrMoneyNegative
	}
	return Money{amount: amount}, nil
}

func MoneyFromStorage(amount int64) Money {
	return Money{amount: amount}
}

func (m Money) Amount() int64                  { return m.amount }
func (m Money) Equals(other Money) bool        { return m.amount == other.amount }
func (m Money) IsGreaterThan(other Money) bool { return m.amount > other.amount }

// PriceSet agrupa las tres invariantes de precio juntas
type PriceSet struct {
	Cost      Money
	Sale      Money
	Wholesale Money
}

func NewPriceSet(cost, sale, wholesale int64) (PriceSet, error) {
	costVO, err := NewMoney(cost)
	if err != nil {
		return PriceSet{}, err
	}
	saleVO, err := NewMoney(sale)
	if err != nil {
		return PriceSet{}, err
	}
	wholesaleVO, err := NewMoney(wholesale)
	if err != nil {
		return PriceSet{}, err
	}

	// invariante 1 — sale_price >= cost_price
	if saleVO.amount < costVO.amount {
		return PriceSet{}, derrors.ErrSalePriceBelowCost
	}
	// invariante 2 — wholesale_price <= sale_price
	if wholesaleVO.amount > saleVO.amount {
		return PriceSet{}, derrors.ErrWholesalePriceAboveSale
	}

	return PriceSet{Cost: costVO, Sale: saleVO, Wholesale: wholesaleVO}, nil
}
