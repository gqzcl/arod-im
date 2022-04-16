package main

import (
	v1 "arod-im/api/connector/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

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

func main() {
	ctx := context.Background()
	//defer cancel()
	ctx.Value(1)
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	client := v1.NewConnectorClient(conn)

	send, err := client.SingleSend(ctx, &v1.SingleSendReq{})
	if err != nil {
		fmt.Println("send failed", err)
	}
	fmt.Println(send.Reply)
}
