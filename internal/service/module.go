package service

import "go.uber.org/fx"

var Module = fx.Module("service",
	fx.Provide(CurrencyServerContainer),
	fx.Invoke(RunService),
)
