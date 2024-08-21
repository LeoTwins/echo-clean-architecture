package handler

import (
	"net/http"
	"strconv"

	"github.com/LeoTwins/go-clean-architecture/internal/interfaces/handler/dto"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/input"
	"github.com/labstack/echo/v4"
)

type ITransactionHandler interface {
	FindByID(c echo.Context) error
	FindByAccountID(c echo.Context) error
}

type transactionHandler struct {
	tu input.ITransactionUsecase
}

func NewTransactionHandler(tu input.ITransactionUsecase) ITransactionHandler {
	return &transactionHandler{tu}
}

func (th *transactionHandler) FindByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	tx, err := th.tu.FindByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res := &dto.TransactionResponse{
		ID:     tx.ID,
		Type:   tx.Type,
		Amount: tx.Amount,
		Date:   tx.Date,
	}

	return c.JSON(http.StatusOK, res)
}

func (th *transactionHandler) FindByAccountID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	transactions, err := th.tu.FindByAccountID(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res := []dto.TransactionResponse{}
	for _, v := range transactions {
		tx := dto.TransactionResponse{
			ID:     v.ID,
			Type:   v.Type,
			Amount: v.Amount,
			Date:   v.Date,
		}

		res = append(res, tx)
	}

	return c.JSON(http.StatusOK, res)
}
