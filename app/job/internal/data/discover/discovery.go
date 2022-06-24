package discover

import (
	ConnectorV1 "arod-im/api/connector/v1"
	"arod-im/pkg/nacos/discover"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

type Discovery struct {
	cli naming_client.INamingClient

	Clients map[string]*ServiceClient
	watcher *discover.Watcher
	//instances []*ServiceInstance

	log *log.Helper
}

// 订阅 -》 chan -》 updateInstance 》 updateClinets

func NewDiscovery(cli naming_client.INamingClient, log *log.Helper) *Discovery {
	c := &Discovery{
		cli: cli,
		log: log,
	}
	c.Clients = make(map[string]*ServiceClient)
	c.watcher, _ = discover.NewWatcher(context.TODO(), c.cli, "arod-im-connector.grpc", "arod-im")
	return c
}

func (c *Discovery) GetClient(address string) ConnectorV1.ConnectorClient {
	// TODO balance
	if client, ok := c.Clients[address]; ok {
		return client.GetClient()
	}
	return nil
}

func (dis *Discovery) Watch() {
	for {
		fmt.Println("正在监听。。。")
		select {
		case <-dis.watcher.Ctx.Done():
			return
		case <-dis.watcher.NoticeChan:
		}
		dis.UpdateClinets(dis.watcher.Instances)
	}

}

func (dis *Discovery) UpdateClinets(ins []*discover.ServiceInstance) {
	// TODO 关闭下线的服务的连接

	// 与新上线的服务建立连接
	for _, instance := range ins {
		addr := instance.Endpoints
		if _, ok := dis.Clients[addr]; !ok {
			var err error
			dis.Clients[addr], err = NewServiceClient(addr)
			if err != nil {
				// 打印
				dis.log.Error(err)
			}
		}
	}
	// fmt.Println("已连接客户端：", dis.Clients)
	dis.log.Debug(dis.Clients)
}

func (dis *Discovery) Close() {
	dis.watcher.Cancel()
}
