package handlers

import (
	proto "Currency-service/pkg/proto"
	"go.uber.org/fx"
)

var Module = fx.Module("handlers",
	fx.Provide(
		CurrencyHandlerConstructor,
		func(h *CurrencyHandler) proto.CurrencyServiceServer { return h },
	),
)
