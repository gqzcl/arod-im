// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package data

// import (
// 	v1 "arod-im/api/connector/v1"
// 	"context"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// 	"google.golang.org/grpc/keepalive"
// 	"time"
// )

// var (
// 	// grpc options
// 	grpcKeepAliveTime    = time.Duration(10) * time.Second
// 	grpcKeepAliveTimeout = time.Duration(3) * time.Second
// 	grpcMaxSendMsgSize   = 1 << 24
// 	grpcMaxCallMsgSize   = 1 << 24
// )

// const (
// 	grpcInitialWindowSize     = 1 << 24
// 	grpcInitialConnWindowSize = 1 << 24
// )

// type ConnectClient struct {
// 	serverID string
// 	client   v1.ConnectorClient
// }

// func NewConnectClient(address string) (*ConnectClient, error) {
// 	cc := &ConnectClient{
// 		serverID: address,
// 	}
// 	cc.client, _ = NewConnectorClient(address)
// 	return cc, nil
// }

// func NewConnectorClient(address string) (v1.ConnectorClient, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*2))
// 	defer cancel()
// 	conn, err := grpc.DialContext(ctx, address,
// 		[]grpc.DialOption{
// 			grpc.WithTransportCredentials(insecure.NewCredentials()),
// 			grpc.WithInitialWindowSize(grpcInitialWindowSize),
// 			grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
// 			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
// 			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
// 			grpc.WithKeepaliveParams(keepalive.ClientParameters{
// 				Time:                grpcKeepAliveTime,
// 				Timeout:             grpcKeepAliveTimeout,
// 				PermitWithoutStream: true,
// 			}),
// 		}...,
// 	)
// 	if err != nil {
// 		panic("grpc 初始化失败")
// 	}
// 	client := v1.NewConnectorClient(conn)
// 	return client, nil
// }

// func (cc *ConnectClient) GetClient() v1.ConnectorClient {
// 	return cc.client
// }
