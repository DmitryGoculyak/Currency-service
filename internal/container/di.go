package container

import (
	"Currency-service/config"
	"Currency-service/internal/db"
	repo "Currency-service/internal/repository/pgsql"
	"Currency-service/internal/service"
	"Currency-service/internal/transport/rpc/handlers"
	"go.uber.org/fx"
)

func Build() *fx.App {
	return fx.New(
		db.Module,
		config.Module,
		service.Module,
		handlers.Module,
		repo.Module,
	)
}
