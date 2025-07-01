package service

import (
	"Currency-service/internal/db/models"
	repo "Currency-service/internal/repository"

	"context"
)

type CurrencyServiceServer interface {
	Create(ctx context.Context, code, name string) error
	GetByCode(ctx context.Context, code string) (*models.CurrencyDB, error)
	ListAll(ctx context.Context) ([]models.CurrencyDB, error)
}

type CurrencyService struct {
	repo repo.CurrencyRepository
}

func CurrencyServiceContainer(
	repo repo.CurrencyRepository,
) *CurrencyService {
	return &CurrencyService{
		repo: repo,
	}
}

func (s *CurrencyService) Create(ctx context.Context, code, name string) error {
	return s.repo.CreateCurrency(ctx, code, name)
}

func (s *CurrencyService) GetByCode(ctx context.Context, code string) (*models.CurrencyDB, error) {
	return s.repo.GetCurrencyByCode(ctx, code)
}

func (s *CurrencyService) ListAll(ctx context.Context) ([]models.CurrencyDB, error) {
	return s.repo.GetAllCurrencies(ctx)
}
