package pgsql

import (
	"Currency-service/internal/entity"
	"context"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CurrencyRepo struct {
	db *sqlx.DB
}

func CurrencyRepoConstructor(
	db *sqlx.DB,
) *CurrencyRepo {
	return &CurrencyRepo{
		db: db,
	}
}

func (r *CurrencyRepo) CreateCurrency(ctx context.Context, code, name string) (*entity.Currency, error) {
	var currency entity.Currency
	err := r.db.GetContext(ctx, &currency, "INSERT INTO currencies(currency_code, currency_name) VALUES ($1, $2)",
		code, name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &currency, nil
}

func (r *CurrencyRepo) GetCurrencyByCode(ctx context.Context, code string) (*entity.Currency, error) {
	var currency entity.Currency
	err := r.db.GetContext(ctx, &currency, "SELECT * FROM currencies WHERE currency_code = $1", code)
	if err != nil {
		return nil, err
	}
	return &currency, nil
}

func (r *CurrencyRepo) GetAllCurrencies(ctx context.Context) ([]entity.Currency, error) {
	var list []entity.Currency
	err := r.db.SelectContext(ctx, &list, "SELECT * FROM currencies")

	return list, err
}
