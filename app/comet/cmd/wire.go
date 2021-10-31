// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"arod-im/app/comet/internal/conf"
	"arod-im/app/comet/internal/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func initApp(*conf.Consul, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, newApp))
}
