package output

import (
	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/dto"
)

type ITransactionPresenter interface {
	Output(model.Transaction) dto.TransactionOutput
}
