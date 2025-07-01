package pgsql

import (
	repo "Currency-service/internal/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("pgsql",
	fx.Provide(
		CurrencyRepoConstructor,
		func(r *CurrencyRepo) repo.CurrencyRepository { return r },
	),
)
