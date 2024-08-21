package output

import (
	"github.com/LeoTwins/go-clean-architecture/internal/domain/model/transaction"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/dto"
)

type ITransactionPresenter interface {
	Output(transaction.Transaction) dto.TransactionOutput
}
