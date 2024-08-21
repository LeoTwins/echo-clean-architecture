package transaction

import (
	"errors"
	"time"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model/finance"
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
	Amount    finance.Money
	Date      time.Time
}

func NewTransaction(id, accountID uint, transactionType TransactionType, amount finance.Money, date time.Time) (*Transaction, error) {
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
