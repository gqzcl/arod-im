package discover

import (
	logicV1 "arod-im/api/logic/v1"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

var (
	// grpc options
	grpcKeepAliveTime    = time.Duration(10) * time.Second
	grpcKeepAliveTimeout = time.Duration(3) * time.Second
	grpcMaxSendMsgSize   = 1 << 24
	grpcMaxCallMsgSize   = 1 << 24
)

const (
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
)

// LogicClient 是logic rpc client
type ServiceClient struct {
	serverID string
	conn     *grpc.ClientConn
	client   logicV1.LogicClient
}

func NewServiceClient(address string) (*ServiceClient, error) {
	lc := &ServiceClient{
		serverID: address,
	}
	err := lc.dailServiceClient(address)
	if err != nil {
		return nil, err
	}
	return lc, nil
}

// create grpc client
func (lc *ServiceClient) dailServiceClient(address string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*2))
	defer cancel()
	conn, err := grpc.DialContext(ctx, address,
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
		return err
	}
	lc.conn = conn
	lc.client = logicV1.NewLogicClient(conn)
	return nil
}

func (sc *ServiceClient) GetClient() logicV1.LogicClient {
	return sc.client
}

func (sc *ServiceClient) Close() {
	sc.conn.Close()
}
