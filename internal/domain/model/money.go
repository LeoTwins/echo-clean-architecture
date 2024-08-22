package model

import "errors"

type Money uint

func NewMoney(amount uint) (*Money, error) {
	if amount <= 0 {
		return nil, errors.New("0円以下は指定できません")
	}
	money := Money(amount)

	return &money, nil
}

func (m Money) Uint() uint {
	return uint(m)
}
