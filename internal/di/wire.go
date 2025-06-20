//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"log/slog"
	"workmate/internal/api"
	"workmate/internal/config"
	somethingProvider "workmate/internal/domain/providers/something"
	taskProvider "workmate/internal/domain/providers/task"
)

func InitializeAPI(cfg *config.Config, log *slog.Logger) (*api.ServerHTTP, error) {
	panic(wire.Build(
		taskProvider.ProviderSet,
		somethingProvider.ProviderSet,
		api.NewServerHTTP))
}
