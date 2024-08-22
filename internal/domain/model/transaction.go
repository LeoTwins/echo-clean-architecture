package model

import (
	"errors"
	"time"
)

type TransactionType string

const (
	Deposit    TransactionType = "DEPOSIT"
	Withdrawal TransactionType = "WITHDRAWAL"
	Transfer   TransactionType = "TRANSFER"
)

type Transaction struct {
	ID        uint
	AccountID uint
	Type      TransactionType
	Amount    Money
	Date      time.Time
}

func NewTransaction(id, accountID uint, transactionType TransactionType, amount Money, date time.Time) (*Transaction, error) {
	if accountID == 0 {
		return nil, errors.New("AccountIDが不正です")
	}

	return &Transaction{
		ID:        id,
		AccountID: accountID,
		Type:      transactionType,
		Amount:    amount,
		Date:      date,
	}, nil
}

func (tt TransactionType) ToString() string {
	return string(tt)
}
