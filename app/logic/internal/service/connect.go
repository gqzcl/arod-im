// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package service

import (
	pb "arod-im/api/logic/v1"
	"context"
)

func (s *MessageService) Connect(ctx context.Context, req *pb.ConnectReq) (*pb.ConnectReply, error) {
	s.log.WithContext(ctx).Debugf("成功收到连接消息form ", req.Uid, req.Address)
	err := s.cc.Connect(ctx, req.Uid, req.Address, req.Server)
	if err != nil {
		return &pb.ConnectReply{
			ActionStatus: "FAIL",
			ErrorInfo:    err.Error(),
			ErrorCode:    90001,
			Success:      false,
		}, err
	}
	return &pb.ConnectReply{
		ActionStatus: "OK",
		Success:      true,
	}, nil
}
func (s *MessageService) Disconnect(ctx context.Context, req *pb.DisConnectReq) (*pb.DisConnectReply, error) {
	s.log.WithContext(ctx).Debugf("成功收到断开连接消息from", req.Uid, req.Address)
	err := s.cc.Disconnect(ctx, req.Uid, req.Address, req.Server)
	if err != nil {
		return &pb.DisConnectReply{
			ActionStatus: "FAIL",
			ErrorInfo:    err.Error(),
			ErrorCode:    90001,
			Success:      false,
		}, err
	}
	return &pb.DisConnectReply{
		ActionStatus: "OK",
		Success:      true,
	}, nil
}
