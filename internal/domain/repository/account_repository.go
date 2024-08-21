package repository

import (
	"context"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model/account"
)

type IAccountRepository interface {
	FindByID(ctx context.Context, id uint) (*account.Account, error)
	Save(ctx context.Context, acc *account.Account) error
	Update(ctx context.Context, acc *account.Account) error
}
