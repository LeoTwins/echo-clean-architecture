package handler

import (
	"net/http"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
	"github.com/LeoTwins/go-clean-architecture/internal/interfaces/handler/dto"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/port/input"
	"github.com/labstack/echo/v4"
)

type IAccountHandler interface {
	OpenAccount(c echo.Context) error
	Deposit(c echo.Context) error
	Withdraw(c echo.Context) error
	Transfer(c echo.Context) error
}

type accountHandler struct {
	au input.IAccountUsecase
}

func NewAccountHandler(au input.IAccountUsecase) IAccountHandler {
	return &accountHandler{au}
}

func (ah *accountHandler) OpenAccount(c echo.Context) error {
	req := dto.OpenAccountRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	money, err := model.NewMoney(req.Balance)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	acc, err := ah.au.OpenAccount(c.Request().Context(), req.Name, *money)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res := dto.OpenAccountResponse{
		ID:      acc.ID,
		Name:    acc.Name,
		Balance: acc.Balance.Value().Uint(),
	}

	return c.JSON(http.StatusCreated, res)
}

func (ah *accountHandler) Deposit(c echo.Context) error {
	req := dto.DepositAndWithdrawRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	money, err := model.NewMoney(req.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := ah.au.Deposit(c.Request().Context(), req.ID, *money); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusOK)
}

func (ah *accountHandler) Withdraw(c echo.Context) error {
	req := dto.DepositAndWithdrawRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	amount, err := model.NewMoney(req.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := ah.au.Withdraw(c.Request().Context(), req.ID, *amount); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusOK)
}

func (ah *accountHandler) Transfer(c echo.Context) error {
	req := dto.TransferRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	money, err := model.NewMoney(req.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err = ah.au.Transfer(c.Request().Context(), req.ID, req.ToAccountID, *money); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusOK)
}
