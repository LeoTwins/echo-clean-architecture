package input

import (
	"context"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
)

type IAccountUsecase interface {
	OpenAccount(ctx context.Context, name string, initialDeposit model.Money) (*model.Account, error)
	Deposit(ctx context.Context, accountID uint, amount model.Money) error
	Withdraw(ctx context.Context, accountID uint, amount model.Money) error
	Transfer(ctx context.Context, fromAccountID uint, toAccountID uint, amount model.Money) error
}
