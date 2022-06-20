// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package service

import (
	pb "arod-im/api/connector/v1"
	logicV1 "arod-im/api/logic/v1"
	"arod-im/app/connector/internal/biz"
	"arod-im/app/connector/internal/conf"
	"arod-im/pkg/ips"
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
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

	bc  *biz.BucketUsecase
	log *log.Helper
	//ws          *websocket.Server
	LogicClient logicV1.LogicClient

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
	ctx := context.Background()
	s.DailLogic(ctx)
	return s
}

func (s *ConnectorService) DailLogic(ctx context.Context) {
	conn, err := grpc.DialContext(ctx, "127.0.0.1:9003",
		[]grpc.DialOption{
			grpc.WithInitialWindowSize(grpcInitialWindowSize),
			grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                grpcKeepAliveTime,
				Timeout:             grpcKeepAliveTimeout,
				PermitWithoutStream: true,
			}),
		}...,
	)
	if err != nil {
		s.log.Error("Grpc 连接失败", err)
	}
	s.LogicClient = logicV1.NewLogicClient(conn)
}

func update(ctx context.Context) {

}
