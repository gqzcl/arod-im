// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package service

import (
	"context"

	pb "arod-im/api/connector/v1"
)

func (s *ConnectorService) SingleSend(ctx context.Context, req *pb.SingleSendReq) (*pb.SingleSendReply, error) {
	s.log.WithContext(ctx).Info("成功收到消息！！！", req.Address, req.Msg)
	s.bc.SingleSend(req.Address, req.Msg)
	return &pb.SingleSendReply{Reply: "成功收到消息！！！"}, nil
}
func (s *ConnectorService) GroupSend(ctx context.Context, req *pb.GroupSendReq) (*pb.GroupSendReply, error) {
	return &pb.GroupSendReply{}, nil
}
func (s *ConnectorService) RoomSend(ctx context.Context, req *pb.RoomSendReq) (*pb.RoomSendReply, error) {
	return &pb.RoomSendReply{}, nil
}
func (s *ConnectorService) Broadcast(ctx context.Context, req *pb.BroadcastReq) (*pb.BroadcastReply, error) {
	return &pb.BroadcastReply{}, nil
}
