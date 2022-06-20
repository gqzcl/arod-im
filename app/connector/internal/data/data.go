// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package data

import (
	"arod-im/app/connector/internal/conf"
	"arod-im/app/connector/internal/data/discover"
	"arod-im/app/connector/internal/data/sender"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBucketRepo)

// Data store resource
type Data struct {
	naming    naming_client.INamingClient
	discovery discover.Discovery

	channel map[string]*sender.Channel
	room    map[string]*sender.Room
	log     *log.Helper
}

// NewData init Data
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	d := &Data{
		// TODO 通过配置文件配置预分配连接数和房间数
		channel: make(map[string]*sender.Channel, 1024),
		room:    make(map[string]*sender.Room, 256),
		log:     log.NewHelper(logger),
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
	d.discovery = *discover.NewDiscovery(d.naming)
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
