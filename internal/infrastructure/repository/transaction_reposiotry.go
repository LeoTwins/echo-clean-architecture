package repository

import (
	"context"
	"errors"

	domain "github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/LeoTwins/go-clean-architecture/internal/domain/repository"
	dbModel "github.com/LeoTwins/go-clean-architecture/internal/infrastructure/repository/model"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func (tr *transactionRepository) FindByAccountID(accountID uint) ([]*domain.Transaction, error) {
	transaction := []dbModel.Transaction{}
	result := tr.db.Where("account_id = ?", accountID).Find(&transaction)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	res := []*domain.Transaction{}
	for _, v := range transaction {
		money, err := domain.NewMoney(v.Amount)
		if err != nil {
			return nil, err
		}

		tx, err := domain.NewTransaction(v.ID, v.AccountID, domain.TransactionType(v.Type), *money, v.Date)
		if err != nil {
			return nil, err
		}

		res = append(res, tx)
	}

	return res, nil
}

func (tr *transactionRepository) FindByID(id uint) (*domain.Transaction, error) {
	transaction := dbModel.Transaction{}
	result := tr.db.Find(&transaction, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	res, err := domain.NewTransaction(transaction.ID, transaction.AccountID, domain.TransactionType(transaction.Type), domain.Money(transaction.Amount), transaction.Date)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (tr *transactionRepository) Save(ctx context.Context, tx *domain.Transaction) error {
	transaction := dbModel.Transaction{
		AccountID: tx.AccountID,
		Amount:    tx.Amount.Uint(),
		Type:      tx.Type.ToString(),
		Date:      tx.Date,
	}

	result := tr.db.Create(&transaction)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewTransactionRepository(db *gorm.DB) repository.ITransactionRepository {
	return &transactionRepository{db}
}
