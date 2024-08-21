package presenter

import (
	"github.com/LeoTwins/go-clean-architecture/internal/domain/model/transaction"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/dto"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/output"
)

type transactionPresenter struct{}

func NewTransactionPresenter() output.ITransactionPresenter {
	return &transactionPresenter{}
}

func (t *transactionPresenter) Output(tx transaction.Transaction) dto.TransactionOutput {
	return dto.TransactionOutput{
		ID:     tx.ID,
		Type:   convertTransactionType(tx.Type),
		Amount: tx.Amount.Uint(),
		Date:   tx.Date.Format("2006/1/2 15:04:05"),
	}
}

func convertTransactionType(t transaction.TransactionType) string {
	switch t {
	case transaction.Deposit:
		return "入金"
	case transaction.Withdrawal:
		return "出金"
	case transaction.Transfer:
		return "振込"
	default:
		return "不明"
	}
}
