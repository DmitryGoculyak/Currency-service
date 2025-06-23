package models

type CurrencyDB struct {
	CurrencyCode string `db:"currency_code"`
	CurrencyName string `db:"currency_name"`
}
