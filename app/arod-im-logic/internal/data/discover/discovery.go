package discover

import (
	"arod-im/pkg/nacos/discover"
	"context"
	"math/rand"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

type Discovery struct {
	cli naming_client.INamingClient

	//Clients map[string]*ServiceClient
	watcher *discover.Watcher
}

// 订阅 -》 chan -》 updateInstance 》 updateClinets

func NewDiscovery(cli naming_client.INamingClient) *Discovery {
	d := &Discovery{
		cli: cli,
	}
	//c.Clients = make(map[string]*ServiceClient)
	d.watcher, _ = discover.NewWatcher(context.TODO(), d.cli, "arod-im-connector.grpc", "arod-im")
	go d.Watch()
	return d
}

// TODO 实现其他负载方式
// GetClient 使用随机负载均衡的方式获取连接服务地址
func (dis *Discovery) GetClient() string {
	lens := len(dis.watcher.Instances)
	if lens == 0 {
		return ""
	}
	rand.Seed(time.Now().Unix())
	return dis.watcher.Instances[rand.Intn(lens)].Endpoints
}

// 将watcher模块的notice chan信号丢弃，避免阻塞
func (dis *Discovery) Watch() {
	for {
		select {
		case <-dis.watcher.Ctx.Done():
			return
		case <-dis.watcher.NoticeChan:
		}
	}
}

func (dis *Discovery) Close() {
	dis.watcher.Cancel()
}
