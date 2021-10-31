package service

import (
	"context"
	"time"

	"arod-im/api/logic/v1"
	"arod-im/api/protocol"
	"arod-im/app/comet/internal/pkg"
	"arod-im/pkg/stringints"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

// Connect connected a connection.
func (s *Server) Connect(c context.Context, p *protocol.Proto, cookie string) (mid int64, key, rid string, accepts []int32, heartbeat time.Duration, err error) {
	reply, err := s.lc.Connect(c, &logic.ConnectReq{
		Server: s.serverID,
		Cookie: cookie,
		Token:  p.Body,
	})
	if err != nil {
		return
	}
	return reply.Mid, reply.Key, reply.RoomID, reply.Accepts, time.Duration(reply.Heartbeat), nil
}

// Disconnect disconnected a connection.
func (s *Server) Disconnect(c context.Context, mid int64, key string) (err error) {
	_, err = s.lc.Disconnect(context.Background(), &logic.DisConnectReq{
		Server: s.serverID,
		Mid:    mid,
		Key:    key,
	})
	return
}

// Heartbeat heartbeat a connection session.
func (s *Server) Heartbeat(ctx context.Context, mid int64, key string) (err error) {
	_, err = s.lc.Heartbeat(ctx, &logic.HeartbeatReq{
		Server: s.serverID,
		Mid:    mid,
		Key:    key,
	})
	return
}

// RenewOnline renew room online.
func (s *Server) RenewOnline(ctx context.Context, serverID string, roomCount map[string]int32) (allRoom map[string]int32, err error) {
	reply, err := s.lc.RenewOnline(ctx, &logic.OnlineReq{
		Server:    s.serverID,
		RoomCount: roomCount,
	}, grpc.UseCompressor(gzip.Name))
	if err != nil {
		return
	}
	return reply.AllRoomCount, nil
}

// Receive receive a message.
func (s *Server) Receive(ctx context.Context, mid int64, p *protocol.Proto) (err error) {
	_, err = s.lc.Receive(ctx, &logic.ReceiveReq{Mid: mid, Proto: p})
	return
}

// Operate operate.
func (s *Server) Operate(ctx context.Context, p *protocol.Proto, ch *pkg.Channel, b *pkg.Bucket) error {
	switch p.Op {
	case protocol.OpChangeRoom:
		if err := b.ChangeRoom(string(p.Body), ch); err != nil {
			glog.Errorf("b.ChangeRoom(%s) error(%v)", p.Body, err)
		}
		p.Op = protocol.OpChangeRoomReply
	case protocol.OpSub:
		// 订阅
		if ops, err := stringints.SplitInt32s(string(p.Body), ","); err == nil {
			ch.Watch(ops...)
		}
		p.Op = protocol.OpSubReply
	case protocol.OpUnsub:
		// 取消订阅
		if ops, err := stringints.SplitInt32s(string(p.Body), ","); err == nil {
			ch.UnWatch(ops...)
		}
		p.Op = protocol.OpUnsubReply
	default:
		// TODO ack ok&failed
		if err := s.Receive(ctx, ch.MemberId, p); err != nil {
			glog.Errorf("s.Report(%d) op:%d error(%v)", ch.MemberId, p.Op, err)
		}
		p.Body = nil
	}
	return nil
}
