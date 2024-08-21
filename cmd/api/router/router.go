package router

import (
	"github.com/LeoTwins/go-clean-architecture/internal/infrastructure/repository"
	"github.com/LeoTwins/go-clean-architecture/internal/infrastructure/service"
	"github.com/LeoTwins/go-clean-architecture/internal/interfaces/handler"
	"github.com/LeoTwins/go-clean-architecture/internal/interfaces/presenter"
	"github.com/LeoTwins/go-clean-architecture/internal/usecase/interacter"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRouter(e *echo.Echo, db *gorm.DB) {
	accountRepo := repository.NewAccountRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	transactionManager := service.NewTransactionManager(db)

	accountUsecase := interacter.NewAccountUsecase(accountRepo, transactionRepo, transactionManager)
	transactionPresenter := presenter.NewTransactionPresenter()
	transactionUsecase := interacter.NewTransactionUsecase(transactionRepo, transactionPresenter)

	accountHandler := handler.NewAccountHandler(accountUsecase)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)

	e.POST("/open-account", accountHandler.OpenAccount)
	e.POST("/deposit", accountHandler.Deposit)
	e.POST("/withdraw", accountHandler.Withdraw)
	e.POST("/transfer", accountHandler.Transfer)

	e.GET("/:id", transactionHandler.FindByID)
	e.GET("/:account-id", transactionHandler.FindByAccountID)
}
