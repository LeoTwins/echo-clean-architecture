package interacter

import (
	"context"
	"time"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/LeoTwins/go-clean-architecture/internal/domain/repository"
	"github.com/LeoTwins/go-clean-architecture/internal/infrastructure/service"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/input"
)

type accountUsecase struct {
	accountRepo        repository.IAccountRepository
	transactionRepo    repository.ITransactionRepository
	transactionManager service.ITransactionManager
}

func NewAccountUsecase(accountRepo repository.IAccountRepository, transactionRepo repository.ITransactionRepository, transactionManager service.ITransactionManager) input.IAccountUsecase {
	return &accountUsecase{accountRepo, transactionRepo, transactionManager}
}

func (a *accountUsecase) OpenAccount(ctx context.Context, name string, initialDeposit model.Money) (*model.Account, error) {
	acc, err := model.NewAccount(0, name, initialDeposit)
	if err != nil {
		return nil, err
	}

	err = a.accountRepo.Save(ctx, acc)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (au *accountUsecase) Deposit(ctx context.Context, accountID uint, amount model.Money) error {
	return au.transactionManager.ExecuteTransaction(func() error {
		acc, err := au.accountRepo.FindByID(ctx, accountID)
		if err != nil {
			return err
		}

		if err = acc.Deposit(amount); err != nil {
			return err
		}

		if err = au.accountRepo.Update(ctx, acc); err != nil {
			return err
		}

		transaction, err := model.NewTransaction(0, accountID, model.Deposit, amount, time.Now())
		if err != nil {
			return err
		}

		if err = au.transactionRepo.Save(ctx, transaction); err != nil {
			return err
		}

		return nil
	})
}

func (au *accountUsecase) Withdraw(ctx context.Context, accountID uint, amount model.Money) error {
	return au.transactionManager.ExecuteTransaction(func() error {
		acc, err := au.accountRepo.FindByID(ctx, accountID)
		if err != nil {
			return err
		}

		if err = acc.WithDraw(amount); err != nil {
			return err
		}

		if err = au.accountRepo.Update(ctx, acc); err != nil {
			return err
		}

		transaction, err := model.NewTransaction(0, accountID, model.Withdrawal, amount, time.Now())
		if err != nil {
			return err
		}

		if err = au.transactionRepo.Save(ctx, transaction); err != nil {
			return err
		}

		return nil
	})
}

func (au *accountUsecase) Transfer(ctx context.Context, fromAccountID uint, toAccountID uint, amount model.Money) error {
	return au.transactionManager.ExecuteTransaction(func() error {
		fromAcc, err := au.accountRepo.FindByID(ctx, fromAccountID)
		if err != nil {
			return err
		}

		toAcc, err := au.accountRepo.FindByID(ctx, toAccountID)
		if err != nil {
			return err
		}

		if err = fromAcc.Transter(toAcc, amount); err != nil {
			return err
		}

		if err = au.accountRepo.Update(ctx, fromAcc); err != nil {
			return err
		}

		if err = au.accountRepo.Update(ctx, toAcc); err != nil {
			return err
		}

		fromTransaction, err := model.NewTransaction(0, fromAccountID, model.Withdrawal, amount, time.Now())
		if err != nil {
			return err
		}

		if err = au.transactionRepo.Save(ctx, fromTransaction); err != nil {
			return err
		}

		toTransaction, err := model.NewTransaction(0, toAccountID, model.Deposit, amount, time.Now())
		if err != nil {
			return err
		}

		if err = au.transactionRepo.Save(ctx, toTransaction); err != nil {
			return err
		}

		return nil
	})
}
