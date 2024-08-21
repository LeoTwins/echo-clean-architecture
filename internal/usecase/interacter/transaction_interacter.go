package interacter

import (
	"github.com/LeoTwins/go-clean-architecture/internal/domain/repository"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/dto"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/input"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/output"
)

type transactionUsecase struct {
	tr        repository.ITransactionRepository
	presenter output.ITransactionPresenter
}

func NewTransactionUsecase(tr repository.ITransactionRepository, presenter output.ITransactionPresenter) input.ITransactionUsecase {
	return &transactionUsecase{tr, presenter}
}

func (tu *transactionUsecase) FindByID(id uint) (*dto.TransactionOutput, error) {
	tx, err := tu.tr.FindByID(id)
	if err != nil {
		return nil, err
	}

	output := tu.presenter.Output(*tx)

	return &output, nil
}

func (tu *transactionUsecase) FindByAccountID(accountID uint) ([]*dto.TransactionOutput, error) {
	tx, err := tu.tr.FindByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	outputs := []*dto.TransactionOutput{}
	for _, v := range tx {
		output := tu.presenter.Output(*v)
		outputs = append(outputs, &output)
	}

	return outputs, nil
}
