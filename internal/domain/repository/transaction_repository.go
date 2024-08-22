package repository

import (
	"context"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
)

type ITransactionRepository interface {
	FindByID(id uint) (*model.Transaction, error)
	FindByAccountID(accountID uint) ([]*model.Transaction, error)
	Save(ctx context.Context, tx *model.Transaction) error
}
