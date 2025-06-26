package container

import (
	"Currency-service/config"
	"Currency-service/internal/db"
	"Currency-service/internal/service"
	"go.uber.org/fx"
)

func Build() *fx.App {
	return fx.New(
		db.Module,
		config.Module,
		service.Module,
	)
}
