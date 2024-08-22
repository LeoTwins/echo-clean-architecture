package repository

import (
	"context"

	"github.com/LeoTwins/go-clean-architecture/internal/domain/model"
)

type IAccountRepository interface {
	FindByID(ctx context.Context, id uint) (*model.Account, error)
	Save(ctx context.Context, acc *model.Account) error
	Update(ctx context.Context, acc *model.Account) error
}
