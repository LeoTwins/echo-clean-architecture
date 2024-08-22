package presenter

import (
	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/dto"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/output"
)

type transactionPresenter struct{}

func NewTransactionPresenter() output.ITransactionPresenter {
	return &transactionPresenter{}
}

func (t *transactionPresenter) Output(tx model.Transaction) dto.TransactionOutput {
	return dto.TransactionOutput{
		ID:     tx.ID,
		Type:   convertTransactionType(tx.Type),
		Amount: tx.Amount.Uint(),
		Date:   tx.Date.Format("2006/1/2 15:04:05"),
	}
}

func convertTransactionType(t model.TransactionType) string {
	switch t {
	case model.Deposit:
		return "入金"
	case model.Withdrawal:
		return "出金"
	case model.Transfer:
		return "振込"
	default:
		return "不明"
	}
}
