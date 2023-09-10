// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package server

import (
	"arod-im/app/arod-im-logic/internal/conf"
	"arod-im/app/arod-im-logic/internal/data"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func NewNacosRegister(c *conf.Server, d *data.Data) *nacos.Registry {
	cc := constant.ClientConfig{
		NamespaceId:         c.Register.NamespaceId,
		AccessKey:           c.Register.AccessKey,
		SecretKey:           c.Register.SecretKey,
		TimeoutMs:           c.Register.TimeoutMs,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            c.Register.LogLevel,
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.Register.Address, c.Register.Port),
	}
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}

	d.SetNaming(namingClient)

	return nacos.New(namingClient, nacos.WithGroup("arod-im"))
}
