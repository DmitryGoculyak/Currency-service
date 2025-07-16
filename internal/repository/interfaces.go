package repository

import (
	"Currency-service/internal/entity"
	"context"
)

type CurrencyRepository interface {
	CreateCurrency(ctx context.Context, code, name string) (*entity.Currency, error)
	GetCurrencyByCode(ctx context.Context, code string) (*entity.Currency, error)
	GetAllCurrencies(ctx context.Context) ([]entity.Currency, error)
}
