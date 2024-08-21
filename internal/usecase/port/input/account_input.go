package input

import (
	"context"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model/account"
	"github.com/LeoTwins/go-clean-architecture/internal/domain/model/finance"
)

type IAccountUsecase interface {
	OpenAccount(ctx context.Context, name string, initialDeposit finance.Money) (*account.Account, error)
	Deposit(ctx context.Context, accountID uint, amount finance.Money) error
	Withdraw(ctx context.Context, accountID uint, amount finance.Money) error
	Transfer(ctx context.Context, fromAccountID uint, toAccountID uint, amount finance.Money) error
}
