package grpc

import (
	"context"
	"net"
	"time"

	pb "github.com/gqzcl/gim/api/comet"
	"github.com/gqzcl/gim/internal/comet"
	"github.com/gqzcl/gim/internal/comet/conf"
	"github.com/gqzcl/gim/internal/comet/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type server struct {
	srv *comet.Server
	pb.UnimplementedCometServer
}

// New comet grpc server.
func New(c *conf.RPCServer, s *comet.Server) *grpc.Server {
	keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
		// MaxConnectionIdle是一个持续时间，表示空闲连接通过发送GoAway关闭的时间量。空闲持续时间是自最近未完成的RPC数变为零或连接建立后定义的。
		MaxConnectionIdle: c.IdleTimeout,
		// MaxConnectionAgeGrace是MaxConnectionAge之后的一个加法时段，在此时段之后，连接将被强制关闭
		MaxConnectionAgeGrace: c.ForceCloseWait,
		// 每隔一段时间ping
		Time: c.KeepAliveInterval,
		// 在ping以进行keepalive检查之后，服务器将等待一段时间的超时，如果在该时间之后仍然没有看到任何活动，则连接将关闭
		Timeout:          c.KeepAliveTimeout,
		MaxConnectionAge: c.MaxLifeTime,
	})
	// NewServer创建一个gRPC服务器，该服务器未注册任何服务，并且尚未开始接受请求
	srv := grpc.NewServer(keepParams)
	pb.RegisterCometServer(srv, &server{s, pb.UnimplementedCometServer{}})
	// 收听本地网络地址上的广播。
	lis, err := net.Listen(c.Network, c.Addr)
	if err != nil {
		panic(err)
	}
	// ？？这里为什么要起一个gorountine
	go func() {
		// Serve 为lis上监听到的每个连接分配一个ServerTransport和service goroutine
		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return srv
}

var _ pb.CometServer = &server{}

// PushMsg push a message to specified sub keys.
func (s *server) PushMsg(ctx context.Context, req *pb.PushMsgReq) (reply *pb.PushMsgReply, err error) {
	if len(req.Keys) == 0 || req.Proto == nil {
		return nil, errors.ErrPushMsgArg
	}
	for _, key := range req.Keys {
		// 根据key从bucket中取得channel
		if channel := s.srv.Bucket(key).Channel(key); channel != nil {
			// 如果不需要发送，继续
			if !channel.NeedPush(req.ProtoOp) {
				continue
			}
			if err = channel.Push(req.Proto); err != nil {
				return
			}
		}
	}
	return &pb.PushMsgReply{}, nil
}

// Broadcast broadcast msg to all user.
func (s *server) Broadcast(ctx context.Context, req *pb.BroadcastReq) (*pb.BroadcastReply, error) {
	if req.Proto == nil {
		return nil, errors.ErrBroadCastArg
	}
	// TODO use broadcast queue
	go func() {
		for _, bucket := range s.srv.Buckets() {
			bucket.Broadcast(req.GetProto(), req.ProtoOp)
			if req.Speed > 0 {
				t := bucket.ChannelCount() / int(req.Speed)
				time.Sleep(time.Duration(t) * time.Second)
			}
		}
	}()
	return &pb.BroadcastReply{}, nil
}

// BroadcastRoom broadcast msg to specified room.
func (s *server) BroadcastRoom(ctx context.Context, req *pb.BroadcastRoomReq) (*pb.BroadcastRoomReply, error) {
	if req.Proto == nil || req.RoomID == "" {
		return nil, errors.ErrBroadCastRoomArg
	}
	for _, bucket := range s.srv.Buckets() {
		bucket.BroadcastRoom(req)
	}
	return &pb.BroadcastRoomReply{}, nil
}

// Rooms gets all the room ids for the server.
func (s *server) Rooms(ctx context.Context, req *pb.RoomsReq) (*pb.RoomsReply, error) {
	var (
		roomIds = make(map[string]bool)
	)
	for _, bucket := range s.srv.Buckets() {
		for roomID := range bucket.Rooms() {
			roomIds[roomID] = true
		}
	}
	return &pb.RoomsReply{Rooms: roomIds}, nil
}
