// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package service

import (
	pb "arod-im/api/connector/v1"
	"arod-im/app/connector/internal/biz"
	"arod-im/app/connector/internal/conf"
	"arod-im/app/connector/internal/service/discover"
	"arod-im/pkg/ips"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewConnectorService)

var (
	// grpc options
	grpcKeepAliveTime    = time.Duration(10) * time.Second
	grpcKeepAliveTimeout = time.Duration(3) * time.Second
	grpcBackoffMaxDelay  = time.Duration(3) * time.Second
	grpcMaxSendMsgSize   = 1 << 24
	grpcMaxCallMsgSize   = 1 << 24
)

const (
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
)

type ConnectorService struct {
	pb.UnimplementedConnectorServer

	naming    naming_client.INamingClient
	discovery discover.Discovery

	bc  *biz.BucketUsecase
	log *log.Helper
	//ws          *websocket.Server
	//LogicClient logicV1.LogicClient

	// server ip
	Address string
}

func NewConnectorService(config *conf.Server, bc *biz.BucketUsecase, logger log.Logger) *ConnectorService {
	ip := ips.InternalIP()
	s := &ConnectorService{
		bc:      bc,
		log:     log.NewHelper(log.With(logger, "module", "connector")),
		Address: fmt.Sprintf("%s:9000", ip),
	}
	return s
}

// SetNaming init the nacos naming client of Data
func (s *ConnectorService) SetNaming(naming naming_client.INamingClient) {
	s.naming = naming
	s.discovery = *discover.NewDiscovery(s.naming)
}

// TODO 如何优雅退出
func (s *ConnectorService) CloseClient() {
	for _, v := range s.discovery.Clients {
		v.Close()
	}
}

func (s *ConnectorService) InitClient() {

	go s.discovery.Watch()
	// s.log.Debugf("已开始监听业务服务地址")
}

// func (s *ConnectorService) DailLogic(ctx context.Context) {
// 	conn, err := grpc.DialContext(ctx, "127.0.0.1:9003",
// 		[]grpc.DialOption{
// 			grpc.WithInitialWindowSize(grpcInitialWindowSize),
// 			grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
// 			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
// 			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
// 			grpc.WithTransportCredentials(insecure.NewCredentials()),
// 			grpc.WithKeepaliveParams(keepalive.ClientParameters{
// 				Time:                grpcKeepAliveTime,
// 				Timeout:             grpcKeepAliveTimeout,
// 				PermitWithoutStream: true,
// 			}),
// 		}...,
// 	)
// 	if err != nil {
// 		s.log.Error("Grpc 连接失败", err)
// 	}
// 	s.LogicClient = logicV1.NewLogicClient(conn)
// }

// func update(ctx context.Context) {

// }
