package repository

import (
	"Currency-service/internal/db/models"
	"context"
)

type CurrencyRepository interface {
	CreateCurrency(ctx context.Context, code, name string) error
	GetCurrencyByCode(ctx context.Context, code string) (*models.CurrencyDB, error)
	GetAllCurrencies(ctx context.Context) ([]models.CurrencyDB, error)
}
