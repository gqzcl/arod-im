package data

import (
	v1 "arod-im/api/connector/v1"
	jobV1 "arod-im/api/job/v1"
	"arod-im/app/job/internal/biz"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"time"
)

var _ biz.JobRepo = (*jobRepo)(nil)

type jobRepo struct {
	data      *Data
	instances *nacos.Registry
	client    v1.ConnectorClient
	log       *log.Helper
}

// NewJobRepo NewGreeterRepo .
func NewJobRepo(data *Data, r *nacos.Registry, logger log.Logger) biz.JobRepo {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "172.20.241.209:9000",
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithInitialWindowSize(grpcInitialWindowSize),
			grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                grpcKeepAliveTime,
				Timeout:             grpcKeepAliveTimeout,
				PermitWithoutStream: true,
			}),
		}...,
	)
	if err != nil {
		panic("grpc 初始化失败")
	}
	client := v1.NewConnectorClient(conn)
	//j.log.WithContext(ctx).Info("成功建立grpc连接")
	return &jobRepo{
		data:      data,
		client:    client,
		instances: r,
		log:       log.NewHelper(logger),
	}
}

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

func (j *jobRepo) SingleSend(ctx context.Context, server, address, seq string, msg []*jobV1.MsgBody) {
	// TODO 处理消息内容
	j.log.Debugf("成功开始发送消息")
	sendReply, err := j.client.SingleSend(ctx, &v1.SingleSendReq{
		Address: address,
		Seq:     seq,
		Msg:     msg,
	})
	if err != nil {
		j.log.WithContext(ctx).Error(err)
	}
	j.log.WithContext(ctx).Info(sendReply)
}
