package data

import (
	v1 "arod-im/api/connector/v1"
	jobV1 "arod-im/api/job/v1"
	"arod-im/app/job/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.JobRepo = (*jobRepo)(nil)

type jobRepo struct {
	data *Data
	log  *log.Helper
}

// NewJobRepo NewGreeterRepo .
func NewJobRepo(data *Data, logger log.Logger) biz.JobRepo {
	//ctx := context.Background()
	//conn, err := grpc.DialContext(ctx, "172.30.105.114:9000",
	//	[]grpc.DialOption{
	//		grpc.WithTransportCredentials(insecure.NewCredentials()),
	//		grpc.WithInitialWindowSize(grpcInitialWindowSize),
	//		grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
	//		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
	//		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
	//		grpc.WithKeepaliveParams(keepalive.ClientParameters{
	//			Time:                grpcKeepAliveTime,
	//			Timeout:             grpcKeepAliveTimeout,
	//			PermitWithoutStream: true,
	//		}),
	//	}...,
	//)
	//if err != nil {
	//	panic("grpc 初始化失败")
	//}
	//_ = v1.NewConnectorClient(conn)
	//j.log.WithContext(ctx).Info("成功建立grpc连接")
	return &jobRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (j *jobRepo) SingleSend(ctx context.Context, server, address, seq string, msg []*jobV1.MsgBody) {
	// TODO 处理消息内容
	j.log.Debugf("成功开始发送消息")
	if connector, ok := j.data.clients[server]; ok {
		client := connector.GetClient()
		sendReply, err := client.SingleSend(ctx, &v1.SingleSendReq{
			Address: address,
			Seq:     seq,
			Msg:     msg,
		})
		if err != nil {
			j.log.WithContext(ctx).Error(err)
		}
		j.log.WithContext(ctx).Info(sendReply)
	} else {
		j.log.WithContext(ctx).Info("Connector 服务地址不存在:", server)
	}

}
