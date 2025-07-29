//go:build wireinject
// +build wireinject

//go:generate wire

package di

import (
	"github.com/google/wire"
)

func SetupApplication() (Application, error) {
	panic(
		wire.Build(
			ProvideContext,
			ProvideLogger,
			ProvideUserRepository,
			ProvideUserService,
			ProvideHTTPServer,
			ProvideGRPCServer,
			NewApplication,
		),
	)
}
