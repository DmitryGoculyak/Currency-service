package service

import (
	"Currency-service/internal/entity"
	repo "Currency-service/internal/repository"

	"context"
)

type CurrencyServiceServer interface {
	Create(ctx context.Context, code, name string) (*entity.Currency, error)
	GetByCode(ctx context.Context, code string) (*entity.Currency, error)
	ListAll(ctx context.Context) ([]entity.Currency, error)
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

func (s *CurrencyService) Create(ctx context.Context, code, name string) (*entity.Currency, error) {
	return s.repo.CreateCurrency(ctx, code, name)
}

func (s *CurrencyService) GetByCode(ctx context.Context, code string) (*entity.Currency, error) {
	return s.repo.GetCurrencyByCode(ctx, code)
}

func (s *CurrencyService) ListAll(ctx context.Context) ([]entity.Currency, error) {
	return s.repo.GetAllCurrencies(ctx)
}
