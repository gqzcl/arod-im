// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package data

import (
	"arod-im/app/arod-im-job/internal/conf"
	"arod-im/app/arod-im-job/internal/data/discover"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewJobRepo)

// Data .
type Data struct {
	naming    naming_client.INamingClient
	discovery discover.Discovery

	log *log.Helper
}

// NewData
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	d := &Data{
		log: log.NewHelper(logger),
	}
	cleanup := func() {
		d.CloseClient()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return d, cleanup, nil
}

// SetNaming init the nacos naming client of Data
func (d *Data) SetNaming(naming naming_client.INamingClient) {
	d.naming = naming
	d.discovery = *discover.NewDiscovery(d.naming, d.log)
}

// TODO 如何优雅退出
func (d *Data) CloseClient() {
	for _, v := range d.discovery.Clients {
		v.Close()
	}
}

func (d *Data) InitClient() {
	go d.discovery.Watch()
}
