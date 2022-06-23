package discover

import (
	logicV1 "arod-im/api/logic/v1"
	"arod-im/pkg/nacos/discover"
	"context"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

type Discovery struct {
	cli naming_client.INamingClient

	Clients map[string]*ServiceClient
	watcher *discover.Watcher
	//instances []*ServiceInstance
}

// 订阅 -》 chan -》 updateInstance 》 updateClinets

func NewDiscovery(cli naming_client.INamingClient) *Discovery {
	c := &Discovery{
		cli: cli,
	}
	c.Clients = make(map[string]*ServiceClient)
	c.watcher, _ = discover.NewWatcher(context.TODO(), c.cli, "arod-im-logic.grpc", "arod-im")
	return c
}

func (c *Discovery) GetClient(address string) logicV1.LogicClient {
	// TODO balance
	return c.Clients[address].GetClient()
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
	for _, instance := range ins {
		addr := instance.Endpoints
		if _, ok := dis.Clients[addr]; !ok {
			var err error
			dis.Clients[addr], err = NewServiceClient(addr)
			if err != nil {
				// 打印
			}
		}
	}
	fmt.Println("已连接客户端：", dis.Clients)
}

func (dis *Discovery) Close() {
	dis.watcher.Cancel()
}
