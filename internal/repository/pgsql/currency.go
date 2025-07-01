package pgsql

import (
	"Currency-service/internal/db/models"
	"context"
	"github.com/jmoiron/sqlx"
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

func (r *CurrencyRepo) CreateCurrency(ctx context.Context, code, name string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO currencies(currency_code, currency_name) VALUES ($1, $2)", code, name)
	return err
}

func (r *CurrencyRepo) GetCurrencyByCode(ctx context.Context, code string) (*models.CurrencyDB, error) {
	var currency models.CurrencyDB
	err := r.db.GetContext(ctx, &currency, "SELECT * FROM currencies WHERE currency_code = $1", code)
	if err != nil {
		return nil, err
	}
	return &currency, nil
}

func (r *CurrencyRepo) GetAllCurrencies(ctx context.Context) ([]models.CurrencyDB, error) {
	var list []models.CurrencyDB
	err := r.db.SelectContext(ctx, &list, "SELECT * FROM currencies")
	return list, err
}
