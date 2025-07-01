package service

import (
	"Currency-service/internal/server"
	"go.uber.org/fx"
)

var Module = fx.Module("service",
	fx.Provide(
		CurrencyServiceContainer,
		func(s *CurrencyService) CurrencyServiceServer { return s },
	),
	fx.Invoke(server.RunService),
)
