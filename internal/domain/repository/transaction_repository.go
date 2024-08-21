package repository

import (
	"context"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model/transaction"
)

type ITransactionRepository interface {
	FindByID(id uint) (*transaction.Transaction, error)
	FindByAccountID(accountID uint) ([]*transaction.Transaction, error)
	Save(ctx context.Context, tx *transaction.Transaction) error
}
