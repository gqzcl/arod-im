// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package data

import (
	"arod-im/app/job/internal/conf"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewJobRepo)

// Data .
type Data struct {
	naming  naming_client.INamingClient
	clients map[string]*ConnectClient

	log *log.Helper
}

// NewData
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	d := &Data{
		clients: make(map[string]*ConnectClient),
		log:     log.NewHelper(logger),
	}
	return d, cleanup, nil
}

func (d *Data) SetNaming(naming naming_client.INamingClient) {
	d.naming = naming
}

func (d *Data) InitClient() {
	d.UpdateInstances()
}

func (d *Data) Watch() {
	d.naming.Subscribe(&vo.SubscribeParam{
		ServiceName: "arod-im-connector.grpc",
		GroupName:   "arod-im",
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			d.UpdateInstances()
		},
	})
}

func (d *Data) UpdateInstances() {
	instances, err := d.naming.SelectInstances(vo.SelectInstancesParam{
		ServiceName: "arod-im-connector.grpc",
		GroupName:   "arod-im",
		HealthyOnly: true,
	})
	if err != nil {
		d.log.Error("获取服务列表失败", err)
	}
	// TODO 新的clinet map
	for _, ins := range instances {
		address := fmt.Sprintf("%s:%d", ins.Ip, ins.Port)
		client, err := NewConnectClient(address)
		if err != nil {
			d.log.Info("grpc 连接失败 in UpdateInstance")
		}
		d.clients[address] = client
		d.log.Info("成功连接grpc with", address)
	}

	fmt.Println("发现所有connector实例", instances)
}
