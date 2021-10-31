package grpc

import (
	"context"
	"net"

	pb "arod-im/api/logic"
	"arod-im/internal/logic"
	"arod-im/internal/logic/conf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type server struct {
	srv *logic.Logic
	// 向前兼容
	pb.UnimplementedLogicServer
}

func New(c *conf.RPCServer, l *logic.Logic) *grpc.Server {
	keepParams := grpc.KeepaliveParams(
		keepalive.ServerParameters{
			MaxConnectionIdle:     c.IdleTimeout,
			MaxConnectionAgeGrace: c.ForceCloseWait,
			Time:                  c.KeepAliveInterval,
			Timeout:               c.KeepAliveTimeout,
			MaxConnectionAge:      c.MaxLifeTime,
		},
	)
	srv := grpc.NewServer(keepParams)
	// ? 向前兼容
	pb.RegisterLogicServer(srv, &server{l, pb.UnimplementedLogicServer{}})
	lis, err := net.Listen(c.Network, c.Addr)
	if err != nil {
		panic(err)
	}
	go func() {
		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return srv
}

var _ pb.LogicServer = &server{}

func (s *server) Connect(ctx context.Context, req *pb.ConnectReq) (*pb.ConnectReply, error) {
	mid, key, room, accepts, hb, err := s.srv.Connect(ctx, req.Server, req.Cookie, req.Token)
	if err != nil {
		return &pb.ConnectReply{}, err
	}
	return &pb.ConnectReply{
		Mid:       mid,
		Key:       key,
		RoomID:    room,
		Accepts:   accepts,
		Heartbeat: hb,
	}, nil
}

// Disconnect disconnect a conn.
func (s *server) Disconnect(ctx context.Context, req *pb.DisConnectReq) (*pb.DisConnectReply, error) {
	has, err := s.srv.Disconnect(ctx, req.Mid, req.Key, req.Server)
	if err != nil {
		return &pb.DisConnectReply{}, err
	}
	return &pb.DisConnectReply{Has: has}, nil
}

// Heartbeat beartbeat a conn.
func (s *server) Heartbeat(ctx context.Context, req *pb.HeartbeatReq) (*pb.HeartbeatReply, error) {
	if err := s.srv.Heartbeat(ctx, req.Mid, req.Key, req.Server); err != nil {
		return &pb.HeartbeatReply{}, err
	}
	return &pb.HeartbeatReply{}, nil
}

// RenewOnline renew server online.
func (s *server) RenewOnline(ctx context.Context, req *pb.OnlineReq) (*pb.OnlineReply, error) {
	allRoomCount, err := s.srv.RenewOnline(ctx, req.Server, req.RoomCount)
	if err != nil {
		return &pb.OnlineReply{}, err
	}
	return &pb.OnlineReply{AllRoomCount: allRoomCount}, nil
}

// Receive receive a message.
func (s *server) Receive(ctx context.Context, req *pb.ReceiveReq) (*pb.ReceiveReply, error) {
	if err := s.srv.Receive(ctx, req.Mid, req.Proto); err != nil {
		return &pb.ReceiveReply{}, err
	}
	return &pb.ReceiveReply{}, nil
}

// nodes return nodes.
func (s *server) Nodes(ctx context.Context, req *pb.NodesReq) (*pb.NodesReply, error) {
	return s.srv.NodesWeighted(ctx, req.Platform, req.ClientIP), nil
}
