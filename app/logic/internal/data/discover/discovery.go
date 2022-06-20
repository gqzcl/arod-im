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
	//instances []*ServiceInstance
}

// 订阅 -》 chan -》 updateInstance 》 updateClinets

func NewDiscovery(cli naming_client.INamingClient) *Discovery {
	c := &Discovery{
		cli: cli,
	}
	//c.Clients = make(map[string]*ServiceClient)
	c.watcher, _ = discover.NewWatcher(context.TODO(), c.cli, "arod-im-connector.grpc", "arod-im")
	return c
}

func (dis *Discovery) GetClient() string {
	// TODO balance
	lens := len(dis.watcher.Instances)
	rand.Seed(time.Now().Unix())
	return dis.watcher.Instances[rand.Intn(lens)].Endpoints
}

// func (dis *Discovery) Watch() {
// 	for {
// 		select {
// 		case <-dis.watcher.Ctx.Done():
// 			return
// 		case <-dis.watcher.WatchChan:
// 		}
// 		dis.UpdateClinets(dis.watcher.Instances)
// 	}

// }

// func (dis *Discovery) UpdateClinets(ins []*discover.ServiceInstance) {
// 	for _, instance := range ins {
// 		addr := instance.Endpoints
// 		if _, ok := dis.Clients[addr]; !ok {
// 			var err error
// 			dis.Clients[addr], err = NewServiceClient(addr)
// 			if err != nil {
// 				// 打印
// 			}
// 		}
// 	}
// }

func (dis *Discovery) Close() {
	dis.watcher.Cancel()
}
