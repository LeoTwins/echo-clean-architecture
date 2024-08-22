package model

import (
	"errors"
	"fmt"
)

const MIN_AMOUNT = 1000

type Account struct {
	ID      uint
	Name    string
	Balance *balance
}

func NewAccount(id uint, name string, amount Money) (*Account, error) {
	if name == "" {
		return nil, errors.New("氏名は必須です")
	}

	balance := NewBalance(amount)

	return &Account{
		ID:      id,
		Name:    name,
		Balance: &balance,
	}, nil
}

func (a *Account) Deposit(amount Money) error {
	if amount < MIN_AMOUNT {
		return errors.New("金額は1,000円以上で指定してください")
	}
	a.Balance.Add(amount)
	return nil
}

func (a *Account) WithDraw(amount Money) error {
	if amount < MIN_AMOUNT {
		return errors.New("金額は1,000円以上で指定してください")
	}
	if err := a.Balance.Subtract(amount); err != nil {
		return fmt.Errorf("出金処理に失敗しました: %v", err)
	}

	return nil
}

func (a *Account) Transter(to *Account, amount Money) error {
	if amount < MIN_AMOUNT {
		return errors.New("金額は1,000円以上で指定してください")
	}

	if to == nil {
		return errors.New("振込先が指定されていません")
	}

	if err := a.Balance.Subtract(amount); err != nil {
		return fmt.Errorf("振り込み処理に失敗しました: %v", err)
	}

	to.Balance.Add(amount)
	return nil
}
