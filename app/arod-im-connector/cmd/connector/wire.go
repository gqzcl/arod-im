// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"arod-im/app/arod-im-connector/internal/biz"
	"arod-im/app/arod-im-connector/internal/conf"
	"arod-im/app/arod-im-connector/internal/data"
	"arod-im/app/arod-im-connector/internal/server"
	"arod-im/app/arod-im-connector/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
