package data

import (
	v1 "arod-im/api/connector/v1"
	"arod-im/app/job/internal/biz"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

var _ biz.JobRepo = (*jobRepo)(nil)

type MsgBody struct {
}

type jobRepo struct {
	data      *Data
	instances *nacos.Registry
	client    v1.ConnectorClient
	log       *log.Helper
}

// NewJobRepo NewGreeterRepo .
func NewJobRepo(data *Data, r *nacos.Registry, logger log.Logger) biz.JobRepo {
	ctx := context.Background()
	conn, _ := grpc.DialContext(ctx, "127.0.0.1:9000",
		[]grpc.DialOption{
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

func (j *jobRepo) SingleSend(ctx context.Context, msg []byte) {
	// TODO 处理消息内容
	sendReply, err := j.client.SingleSend(ctx, &v1.SingleSendReq{
		Address: "",
		Msg:     nil,
	})
	if err != nil {
		return
	}
	j.log.WithContext(ctx).Info(sendReply)
}

func (j *jobRepo) Consumer(ctx context.Context) error {
	m, err := j.data.consumer.ReadMessage(ctx)
	if err != nil {
		j.log.WithContext(ctx).Error("err in data Consumer", err)
		return err
	}
	// TODO 将消息送到对应的server
	//_, err = j.instances.GetService(ctx, "arod-im-connector.grpc")
	//if err != nil {
	//	j.log.Error(" 服务发现失败")
	//}

	send, err := j.client.SingleSend(ctx, &v1.SingleSendReq{Address: "127.0.0.1:38478", Msg: m.Value})
	if err != nil {
		j.log.WithContext(ctx).Error("err in data SingleSend", err)
		return err
	}
	j.log.WithContext(ctx).Info("返回响应", send.Reply)

	j.log.WithContext(ctx).Infof("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	return nil
}
