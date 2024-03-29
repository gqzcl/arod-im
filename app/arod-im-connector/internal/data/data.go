// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package data

import (
	"arod-im/app/arod-im-connector/internal/conf"
	"arod-im/app/arod-im-connector/internal/data/sender"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBucketRepo)

// Data store resource
type Data struct {
	//naming    naming_client.INamingClient
	//discovery discover.Discovery

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
		log.NewHelper(logger).Info("closing the data resources")
	}
	return d, cleanup, nil
}
