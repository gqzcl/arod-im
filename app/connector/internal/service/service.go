package service

import (
	pb "arod-im/api/connector/v1"
	"arod-im/app/connector/internal/biz"
	"arod-im/pkg/transport/websocket"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewConnectorService)

type ConnectorService struct {
	pb.UnimplementedConnectorServer
	ws  *websocket.Server
	bc  *biz.BucketUsecase
	log *log.Helper
}

func NewConnectorService(bc *biz.BucketUsecase, logger log.Logger) *ConnectorService {
	return &ConnectorService{
		bc:  bc,
		log: log.NewHelper(log.With(logger, "module", "connector")),
	}
}
