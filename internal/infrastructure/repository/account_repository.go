package repository

import (
	"context"
	"errors"
	"log"

	domain "github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/LeoTwins/go-clean-architecture/internal/domain/repository"
	dbModel "github.com/LeoTwins/go-clean-architecture/internal/infrastructure/repository/model"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) repository.IAccountRepository {
	return &accountRepository{db}
}

func (ar *accountRepository) FindByID(ctx context.Context, id uint) (*domain.Account, error) {
	var dbAcc dbModel.Account
	result := ar.db.First(&dbAcc, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	money, err := domain.NewMoney(dbAcc.Balance)
	if err != nil {
		return nil, err
	}

	acc, err := domain.NewAccount(dbAcc.ID, dbAcc.Name, *money)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (ar *accountRepository) Save(ctx context.Context, acc *domain.Account) error {
	dbAcc := dbModel.Account{
		Name:    acc.Name,
		Balance: acc.Balance.Value().Uint(),
	}

	result := ar.db.Create(&dbAcc)
	if result.Error != nil {
		return result.Error
	}
	acc.ID = dbAcc.ID
	return nil
}

func (ar *accountRepository) Update(ctx context.Context, acc *domain.Account) error {
	log.Print(acc)
	dbAcc := dbModel.Account{
		ID:      acc.ID,
		Name:    acc.Name,
		Balance: acc.Balance.Value().Uint(),
	}

	result := ar.db.Save(&dbAcc)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
