package service

import (
	pb "arod-im/api/logic/v1"
	"context"
)

func (s *MessageService) Connect(ctx context.Context, req *pb.ConnectReq) (*pb.ConnectReply, error) {
	s.log.WithContext(ctx).Info("成功收到连接消息！！！")
	s.cc.Connect(ctx, req.Uid, req.Address, req.Server)
	return &pb.ConnectReply{Success: true}, nil
}
func (s *MessageService) Disconnect(ctx context.Context, req *pb.DisConnectReq) (*pb.DisConnectReply, error) {
	s.log.WithContext(ctx).Info("成功收到断开连接消息！！！")
	s.cc.Disconnect(ctx, req.Uid, req.Address, req.Server)
	return &pb.DisConnectReply{Success: true}, nil
}
