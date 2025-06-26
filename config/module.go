package config

import (
	"Currency-service/internal/db"
	"go.uber.org/fx"
)

var Module = fx.Module("config",
	fx.Provide(
		LoadConfig,
		func(cfg *Config) db.DBConfig { return *cfg.DBConfig },
		func(cfg *Config) *GrpcServiceConfig { return cfg.GrpcConfig },
	),
)
