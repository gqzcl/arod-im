package discover

import (
	"context"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type ServiceInstance struct {
	// ID is the unique instance ID as registered.
	ID string `json:"id"`
	// Name is the service name as registered.
	Name string `json:"name"`
	// Version is the version of the compiled.
	Version string `json:"version"`
	// Metadata is the kv pair metadata associated with the service instance.
	Metadata map[string]string `json:"metadata"`
	// Endpoints is endpoint addresses of the service instance.
	// schema:
	//   http://127.0.0.1:8000?isSecure=false
	//   grpc://127.0.0.1:9000?isSecure=false
	Endpoints string `json:"endpoints"`
}
type Watcher struct {
	serviceName string
	clusters    []string
	groupName   string
	Ctx         context.Context
	Cancel      context.CancelFunc
	watchChan   chan struct{}
	NoticeChan  chan struct{}
	cli         naming_client.INamingClient
	kind        string
	Instances   []*ServiceInstance
}

type options func(w *Watcher)

func NewWatcher(ctx context.Context, cli naming_client.INamingClient, serviceName, groupName string, opts ...options) (*Watcher, error) {
	w := &Watcher{
		serviceName: serviceName,
		cli:         cli,
		groupName:   groupName,
		watchChan:   make(chan struct{}, 1),
		NoticeChan:  make(chan struct{}, 1),
	}
	for _, opt := range opts {
		opt(w)
	}
	w.GetService()
	w.Ctx, w.Cancel = context.WithCancel(ctx)
	e := w.cli.Subscribe(&vo.SubscribeParam{
		ServiceName: serviceName,
		Clusters:    w.clusters,
		GroupName:   w.groupName,
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			//TODO bug 当监听服务集群为空时，不会触发回调函数
			w.watchChan <- struct{}{}
		},
	})
	go w.Update()
	return w, e
}

func (w *Watcher) Update() {
	for {
		select {
		case <-w.Ctx.Done():
			return
		case <-w.watchChan:
		}
		err := w.GetService()
		if err != nil {
			return
		}
		w.NoticeChan <- struct{}{}
	}
}

func (w *Watcher) GetService() error {
	// TODO 加锁
	res, err := w.cli.GetService(vo.GetServiceParam{
		ServiceName: w.serviceName,
		GroupName:   w.groupName,
		Clusters:    w.clusters,
	})
	if err != nil {
		return err
	}
	w.Instances = make([]*ServiceInstance, 0, len(res.Hosts))
	for _, in := range res.Hosts {
		// kind := w.kind
		// if k, ok := in.Metadata["kind"]; ok {
		// 	kind = k
		// }
		w.Instances = append(w.Instances, &ServiceInstance{
			ID:       in.InstanceId,
			Name:     res.Name,
			Version:  in.Metadata["version"],
			Metadata: in.Metadata,
			// Endpoints: fmt.Sprintf("%s://%s:%d", kind, in.Ip, in.Port),
			Endpoints: fmt.Sprintf("%s:%d", in.Ip, in.Port),
		})
	}
	return nil
}

func (w *Watcher) Stop() error {
	w.Cancel()
	return w.cli.Unsubscribe(&vo.SubscribeParam{
		ServiceName: w.serviceName,
		GroupName:   w.groupName,
		Clusters:    w.clusters,
	})
}

// Options
func WithCluster(clusters []string) options {
	return func(w *Watcher) {
		w.clusters = clusters
	}
}

func WithKind(kind string) options {
	return func(w *Watcher) {
		w.kind = kind
	}
}
