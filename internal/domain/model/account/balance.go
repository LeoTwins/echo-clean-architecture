package account

import (
	"errors"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model/finance"
)

type balance uint

func NewBalance(amount finance.Money) balance {
	return balance(amount)
}

func (b *balance) Value() finance.Money {
	return finance.Money(*b)
}

func (b *balance) Add(amount finance.Money) {
	*b += NewBalance(amount)
}

func (b *balance) Subtract(amount finance.Money) error {
	if b.Value() < amount {
		return errors.New("預金残高が不足しています")
	}
	*b -= NewBalance(amount)
	return nil
}
