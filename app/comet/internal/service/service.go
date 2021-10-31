package service

import (
	v1 "arod-im/api/comet/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewCometService)

type CometService struct {
	v1.UnimplementedCometServer
	log *log.Helper
}
type cometServiceOption func()

func InitTCP(addrs string, logger log.Logger) cometServiceOption {
	return func() {
		InitTCPServer(&Server{}, addrs, logger)
	}
}
func InitWebsocket(addrs string, logger log.Logger) cometServiceOption {
	return func() {
		InitWebsocketServer(&Server{}, addrs, logger)
	}
}
func CometServiceOptions(opts ...cometServiceOption) cometServiceOption {
	return func() {
		for _, opt := range opts {
			opt()
		}
	}
}

func NewCometService(logger log.Logger, opts ...cometServiceOption) *CometService {
	for _, opt := range opts {
		opt()
	}
	return &CometService{
		log: log.NewHelper(log.With(logger, "module", "comet/service")),
	}
}
