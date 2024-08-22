package model

import (
	"errors"
)

type balance uint

func NewBalance(amount Money) balance {
	return balance(amount)
}

func (b *balance) Value() Money {
	return Money(*b)
}

func (b *balance) Add(amount Money) {
	*b += NewBalance(amount)
}

func (b *balance) Subtract(amount Money) error {
	if b.Value() < amount {
		return errors.New("預金残高が不足しています")
	}
	*b -= NewBalance(amount)
	return nil
}
